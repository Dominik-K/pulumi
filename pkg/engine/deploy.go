// Copyright 2017, Pulumi Corporation.  All rights reserved.

package engine

import (
	"bytes"
	"fmt"
	"time"

	goerr "github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/compiler/errors"
	"github.com/pulumi/pulumi/pkg/diag"
	"github.com/pulumi/pulumi/pkg/diag/colors"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/resource/deploy"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/pulumi/pulumi/pkg/util/contract"
)

type DeployOptions struct {
	Package              string   // the package we are deploying (or "" to use the default)
	Analyzers            []string // an optional set of analyzers to run as part of this deployment.
	Debug                bool     // true to enable resource debugging output.
	DryRun               bool     // true if we should just print the plan without performing it.
	Parallel             int      // the degree of parallelism for resource operations (<=1 for serial).
	ShowConfig           bool     // true to show the configuration variables being used.
	ShowReplacementSteps bool     // true to show the replacement steps in the plan.
	ShowSames            bool     // true to show the resources that aren't updated, in addition to those that are.
	Summary              bool     // true if we should only summarize resources and operations.
}

func (eng *Engine) Deploy(environment tokens.QName, opts DeployOptions) error {
	contract.Require(environment != tokens.QName(""), "environment")

	// Initialize the diagnostics logger with the right stuff.
	eng.InitDiag(diag.FormatOptions{
		Colors: true,
		Debug:  opts.Debug,
	})

	info, err := eng.planContextFromEnvironment(environment, opts.Package)
	if err != nil {
		return err
	}
	return eng.deployLatest(info, deployOptions{
		Debug:                opts.Debug,
		Destroy:              false,
		DryRun:               opts.DryRun,
		Analyzers:            opts.Analyzers,
		Parallel:             opts.Parallel,
		ShowConfig:           opts.ShowConfig,
		ShowReplacementSteps: opts.ShowReplacementSteps,
		ShowSames:            opts.ShowSames,
		Summary:              opts.Summary,
	})
}

type deployOptions struct {
	Debug                bool     // true to enable resource debugging output.
	Create               bool     // true if we are creating resources.
	Destroy              bool     // true if we are destroying the environment.
	DryRun               bool     // true if we should just print the plan without performing it.
	Analyzers            []string // an optional set of analyzers to run as part of this deployment.
	Parallel             int      // the degree of parallelism for resource operations (<=1 for serial).
	ShowConfig           bool     // true to show the configuration variables being used.
	ShowReplacementSteps bool     // true to show the replacement steps in the plan.
	ShowSames            bool     // true to show the resources that aren't updated, in addition to those that are.
	Summary              bool     // true if we should only summarize resources and operations.
	DOT                  bool     // true if we should print the DOT file for this plan.
}

func (eng *Engine) deployLatest(info *planContext, opts deployOptions) error {
	result, err := eng.plan(info, opts)
	if err != nil {
		return err
	}
	if result != nil {
		defer contract.IgnoreClose(result)
		if opts.DryRun {
			// If a dry run, just print the plan, don't actually carry out the deployment.
			if err := eng.printPlan(result); err != nil {
				return err
			}
		} else {
			// Otherwise, we will actually deploy the latest bits.
			var header bytes.Buffer
			printPrelude(&header, result, false)
			header.WriteString(fmt.Sprintf("%vPerforming changes:%v\n", colors.SpecUnimportant, colors.Reset))
			fmt.Fprint(eng.Stdout, colors.Colorize(&header))

			// Walk the plan, reporting progress and executing the actual operations as we go.
			start := time.Now()
			actions := &deployActions{
				Ops:    make(map[deploy.StepOp]int),
				Opts:   opts,
				Target: result.Info.Target,
				Engine: eng,
			}
			summary, _, _, err := result.Walk(actions)
			if err != nil && summary == nil {
				// Something went wrong, and no changes were made.
				return err
			}
			contract.Assert(summary != nil)

			// Print a summary.
			var footer bytes.Buffer
			// Print out the total number of steps performed (and their kinds), the duration, and any summary info.
			if c := printChangeSummary(&footer, actions.Ops, false); c != 0 {
				footer.WriteString(fmt.Sprintf("%vUpdate duration: %v%v\n",
					colors.SpecUnimportant, time.Since(start), colors.Reset))
			}

			if actions.MaybeCorrupt {
				footer.WriteString(fmt.Sprintf(
					"%vA catastrophic error occurred; resources states may be unknown%v\n",
					colors.SpecAttention, colors.Reset))
			}

			fmt.Fprint(eng.Stdout, colors.Colorize(&footer))
			if err != nil {
				return err
			}
		}
	}
	if !eng.Diag().Success() {
		// If any error that wasn't printed above, be sure to make it evident in the output.
		return goerr.New("One or more errors occurred during this update")
	}
	return nil
}

// deployActions pretty-prints the plan application process as it goes.
type deployActions struct {
	Steps        int
	Ops          map[deploy.StepOp]int
	MaybeCorrupt bool
	Opts         deployOptions
	Target       *deploy.Target
	Engine       *Engine
}

func (acts *deployActions) Run(step deploy.Step) (resource.Status, error) {
	// Report the beginning of the step if appropriate.
	if shouldShow(step, acts.Opts) {
		var b bytes.Buffer
		printStep(&b, step, acts.Opts.Summary, false, "")
		fmt.Fprint(acts.Engine.Stdout, colors.Colorize(&b))
	}

	// Apply the step's changes.
	status, err := step.Apply()

	// Report the result of the step.
	stepop := step.Op()
	if err != nil {
		// Issue a true, bonafide error.
		acts.Engine.Diag().Errorf(errors.ErrorPlanApplyFailed, err)

		// Print the state of the resource; we don't issue the error, because the deploy above will do that.
		var b bytes.Buffer
		stepnum := acts.Steps + 1
		b.WriteString(fmt.Sprintf("Step #%v failed [%v]: ", stepnum, stepop))
		switch status {
		case resource.StatusOK:
			b.WriteString(colors.SpecNote)
			b.WriteString("provider successfully recovered from this failure")
		case resource.StatusUnknown:
			b.WriteString(colors.SpecAttention)
			b.WriteString("this failure was catastrophic and the provider cannot guarantee recovery")
			acts.MaybeCorrupt = true
		default:
			contract.Failf("Unrecognized resource state: %v", status)
		}
		b.WriteString(colors.Reset)
		b.WriteString("\n")
		fmt.Fprint(acts.Engine.Stdout, colors.Colorize(&b))
	} else {
		// Increment the counters.
		if step.Logical() {
			acts.Steps++
			acts.Ops[stepop]++
		}

		// Print out any output properties that got created as a result of this operation.
		if shouldShow(step, acts.Opts) && !acts.Opts.Summary {
			var b bytes.Buffer
			printResourceOutputProperties(&b, step, "")
			fmt.Fprint(acts.Engine.Stdout, colors.Colorize(&b))
		}
	}

	// Write out the current snapshot. Note that even if a failure has occurred, we should still have
	// a safe checkpoint. Note that any error that occurs when writing the checkpoint trumps the error
	// reported above.
	saveErr := acts.Engine.Environment.SaveEnvironment(acts.Target, step.Iterator().Snap())
	if saveErr != nil {
		err = saveErr
	}

	return status, err
}
