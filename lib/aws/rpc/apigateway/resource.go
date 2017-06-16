// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/resource/plugin"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for Resource resource provider */

// ResourceToken is the type token corresponding to the Resource package type.
const ResourceToken = tokens.Type("aws:apigateway/resource:Resource")

// ResourceProviderOps is a pluggable interface for Resource-related management functionality.
type ResourceProviderOps interface {
    Check(ctx context.Context, obj *Resource, property string) error
    Create(ctx context.Context, obj *Resource) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Resource, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Resource, new *Resource, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Resource, new *Resource, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// ResourceProvider is a dynamic gRPC-based plugin for managing Resource resources.
type ResourceProvider struct {
    ops ResourceProviderOps
}

// NewResourceProvider allocates a resource provider that delegates to a ops instance.
func NewResourceProvider(ops ResourceProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &ResourceProvider{ops: ops}
}

func (p *ResourceProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Resource", "name", failure))
        }
    }
    if !unks["parent"] {
        if failure := p.ops.Check(ctx, obj, "parent"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Resource", "parent", failure))
        }
    }
    if !unks["pathPart"] {
        if failure := p.ops.Check(ctx, obj, "pathPart"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Resource", "pathPart", failure))
        }
    }
    if !unks["restAPI"] {
        if failure := p.ops.Check(ctx, obj, "restAPI"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Resource", "restAPI", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *ResourceProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Resource_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *ResourceProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *ResourceProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: plugin.MarshalProperties(
            nil, resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *ResourceProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("parent") {
            replaces = append(replaces, "parent")
        }
        if diff.Changed("pathPart") {
            replaces = append(replaces, "pathPart")
        }
        if diff.Changed("restAPI") {
            replaces = append(replaces, "restAPI")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *ResourceProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ResourceProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ResourceToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ResourceProvider) Unmarshal(
    v *pbstruct.Struct) (*Resource, resource.PropertyMap, error) {
    var obj Resource
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Resource structure(s) */

// Resource is a marshalable representation of its corresponding IDL type.
type Resource struct {
    Name *string `lumi:"name,optional"`
    Parent resource.ID `lumi:"parent"`
    PathPart string `lumi:"pathPart"`
    RestAPI resource.ID `lumi:"restAPI"`
}

// Resource's properties have constants to make dealing with diffs and property bags easier.
const (
    Resource_Name = "name"
    Resource_Parent = "parent"
    Resource_PathPart = "pathPart"
    Resource_RestAPI = "restAPI"
)


