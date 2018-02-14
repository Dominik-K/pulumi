// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/pulumi/pulumi/pkg/workspace"
)

func TestPrettyKeyForProject(t *testing.T) {
	proj := &workspace.Project{Name: tokens.PackageName("test-package"), Runtime: "nodejs"}
	assert.Equal(t, "foo", prettyKeyForProject("test-package:config:foo", proj))
	assert.Equal(t, "other-package:config:bar", prettyKeyForProject("other-package:config:bar", proj))
}
