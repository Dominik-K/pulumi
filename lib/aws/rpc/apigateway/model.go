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

/* RPC stubs for Model resource provider */

// ModelToken is the type token corresponding to the Model package type.
const ModelToken = tokens.Type("aws:apigateway/model:Model")

// ModelProviderOps is a pluggable interface for Model-related management functionality.
type ModelProviderOps interface {
    Check(ctx context.Context, obj *Model, property string) error
    Create(ctx context.Context, obj *Model) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Model, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Model, new *Model, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Model, new *Model, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// ModelProvider is a dynamic gRPC-based plugin for managing Model resources.
type ModelProvider struct {
    ops ModelProviderOps
}

// NewModelProvider allocates a resource provider that delegates to a ops instance.
func NewModelProvider(ops ModelProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &ModelProvider{ops: ops}
}

func (p *ModelProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "name", failure))
        }
    }
    if !unks["contentType"] {
        if failure := p.ops.Check(ctx, obj, "contentType"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "contentType", failure))
        }
    }
    if !unks["restAPI"] {
        if failure := p.ops.Check(ctx, obj, "restAPI"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "restAPI", failure))
        }
    }
    if !unks["schema"] {
        if failure := p.ops.Check(ctx, obj, "schema"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "schema", failure))
        }
    }
    if !unks["modelName"] {
        if failure := p.ops.Check(ctx, obj, "modelName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "modelName", failure))
        }
    }
    if !unks["description"] {
        if failure := p.ops.Check(ctx, obj, "description"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Model", "description", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *ModelProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Model_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *ModelProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
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

func (p *ModelProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
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

func (p *ModelProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
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
        if diff.Changed("contentType") {
            replaces = append(replaces, "contentType")
        }
        if diff.Changed("restAPI") {
            replaces = append(replaces, "restAPI")
        }
        if diff.Changed("modelName") {
            replaces = append(replaces, "modelName")
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

func (p *ModelProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ModelToken))
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

func (p *ModelProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ModelProvider) Unmarshal(
    v *pbstruct.Struct) (*Model, resource.PropertyMap, error) {
    var obj Model
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Model structure(s) */

// Model is a marshalable representation of its corresponding IDL type.
type Model struct {
    Name *string `lumi:"name,optional"`
    ContentType string `lumi:"contentType"`
    RestAPI resource.ID `lumi:"restAPI"`
    Schema interface{} `lumi:"schema"`
    ModelName *string `lumi:"modelName,optional"`
    Description *string `lumi:"description,optional"`
}

// Model's properties have constants to make dealing with diffs and property bags easier.
const (
    Model_Name = "name"
    Model_ContentType = "contentType"
    Model_RestAPI = "restAPI"
    Model_Schema = "schema"
    Model_ModelName = "modelName"
    Model_Description = "description"
)


