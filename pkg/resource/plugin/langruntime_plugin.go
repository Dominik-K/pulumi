// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package plugin

import (
	"strings"

	"github.com/blang/semver"
	"github.com/golang/glog"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/pulumi/pulumi/pkg/util/contract"
	"github.com/pulumi/pulumi/pkg/workspace"
	pulumirpc "github.com/pulumi/pulumi/sdk/proto/go"
)

// langhost reflects a language host plugin, loaded dynamically for a single language/runtime pair.
type langhost struct {
	ctx     *Context
	runtime string
	plug    *plugin
	client  pulumirpc.LanguageRuntimeClient
}

// NewLanguageRuntime binds to a language's runtime plugin and then creates a gRPC connection to it.  If the
// plugin could not be found, or an error occurs while creating the child process, an error is returned.
func NewLanguageRuntime(host Host, ctx *Context, runtime string) (LanguageRuntime, error) {
	// Load the plugin's path by using the standard workspace logic.
	_, path, err := workspace.GetPluginPath(
		workspace.LanguagePlugin, strings.Replace(runtime, tokens.QNameDelimiter, "_", -1), nil)
	if err != nil {
		return nil, err
	} else if path == "" {
		return nil, NewMissingError(workspace.PluginInfo{
			Kind: workspace.LanguagePlugin,
			Name: runtime,
		})
	}

	plug, err := newPlugin(ctx, path, runtime, []string{host.ServerAddr()})
	if err != nil {
		return nil, err
	}
	contract.Assertf(plug != nil, "unexpected nil language plugin for %s", runtime)

	return &langhost{
		ctx:     ctx,
		runtime: runtime,
		plug:    plug,
		client:  pulumirpc.NewLanguageRuntimeClient(plug.Conn),
	}, nil
}

func (h *langhost) Runtime() string { return h.runtime }

// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
func (h *langhost) GetRequiredPlugins(info ProgInfo) ([]workspace.PluginInfo, error) {
	proj := string(info.Proj.Name)
	glog.V(7).Infof("langhost[%v].GetRequiredPlugins(proj=%s,pwd=%s,program=%s) executing",
		h.runtime, proj, info.Pwd, info.Program)
	resp, err := h.client.GetRequiredPlugins(h.ctx.Request(), &pulumirpc.GetRequiredPluginsRequest{
		Project: proj,
		Pwd:     info.Pwd,
		Program: info.Program,
	})
	if err != nil {
		glog.V(7).Infof("langhost[%v].GetRequiredPlugins(proj=%s,pwd=%s,program=%s) failed: err=%v",
			h.runtime, proj, info.Pwd, info.Program, err)

		// It's possible this is just an older language host, prior to the emergence of the GetRequiredPlugins
		// method.  In such cases, we will silently error (with the above log left behind).
		if staterr, ok := status.FromError(err); ok && staterr.Code() == codes.Unimplemented {
			return nil, nil
		}

		return nil, err
	}

	var results []workspace.PluginInfo
	for _, info := range resp.GetPlugins() {
		var version *semver.Version
		if v := info.GetVersion(); v != "" {
			sv, err := semver.ParseTolerant(v)
			if err != nil {
				return nil, errors.Wrapf(err, "illegal semver returned by language host: %s@%s", info.GetName(), v)
			}
			version = &sv
		}
		if !workspace.IsPluginKind(info.GetKind()) {
			return nil, errors.Errorf("unrecognized plugin kind: %s", info.GetKind())
		}
		results = append(results, workspace.PluginInfo{
			Name:    info.GetName(),
			Kind:    workspace.PluginKind(info.GetKind()),
			Version: version,
		})
	}

	glog.V(7).Infof("langhost[%v].GetRequiredPlugins(proj=%s,pwd=%s,program=%s) success: #versions=%d",
		h.runtime, proj, info.Pwd, info.Program, len(results))
	return results, nil

}

// Run executes a program in the language runtime for planning or deployment purposes.  If info.DryRun is true,
// the code must not assume that side-effects or final values resulting from resource deployments are actually
// available.  If it is false, on the other hand, a real deployment is occurring and it may safely depend on these.
func (h *langhost) Run(info RunInfo) (string, error) {
	glog.V(7).Infof("langhost[%v].Run(pwd=%v,program=%v,#args=%v,proj=%s,stack=%v,#config=%v,dryrun=%v) executing",
		h.runtime, info.Pwd, info.Program, len(info.Args), info.Project, info.Stack, len(info.Config), info.DryRun)
	config := make(map[string]string)
	for k, v := range info.Config {
		config[string(k.AsModuleMember())] = v
	}
	resp, err := h.client.Run(h.ctx.Request(), &pulumirpc.RunRequest{
		MonitorAddress: info.MonitorAddress,
		Pwd:            info.Pwd,
		Program:        info.Program,
		Args:           info.Args,
		Project:        info.Project,
		Stack:          info.Stack,
		Config:         config,
		DryRun:         info.DryRun,
		Parallel:       int32(info.Parallel),
	})
	if err != nil {
		glog.V(7).Infof("langhost[%v].Run(pwd=%v,program=%v,...,dryrun=%v) failed: err=%v",
			h.runtime, info.Pwd, info.Program, info.DryRun, err)
		return "", err
	}

	progerr := resp.GetError()
	glog.V(7).Infof("langhost[%v].RunPlan(pwd=%v,program=%v,...,dryrun=%v) success: progerr=%v",
		h.runtime, info.Pwd, info.Program, info.DryRun, progerr)
	return progerr, nil
}

// GetPluginInfo returns this plugin's information.
func (h *langhost) GetPluginInfo() (workspace.PluginInfo, error) {
	glog.V(7).Infof("langhost[%v].GetPluginInfo() executing", h.runtime)
	resp, err := h.client.GetPluginInfo(h.ctx.Request(), &pbempty.Empty{})
	if err != nil {
		glog.V(7).Infof("langhost[%v].GetPluginInfo() failed: err=%v", h.runtime, err)
		return workspace.PluginInfo{}, err
	}

	var version *semver.Version
	if v := resp.Version; v != "" {
		sv, err := semver.ParseTolerant(v)
		if err != nil {
			return workspace.PluginInfo{}, err
		}
		version = &sv
	}

	return workspace.PluginInfo{
		Name:    h.runtime,
		Path:    h.plug.Bin,
		Kind:    workspace.LanguagePlugin,
		Version: version,
	}, nil
}

// Close tears down the underlying plugin RPC connection and process.
func (h *langhost) Close() error {
	return h.plug.Close()
}
