// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

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

/* RPC stubs for Bucket resource provider */

// BucketToken is the type token corresponding to the Bucket package type.
const BucketToken = tokens.Type("aws:s3/bucket:Bucket")

// BucketProviderOps is a pluggable interface for Bucket-related management functionality.
type BucketProviderOps interface {
    Check(ctx context.Context, obj *Bucket, property string) error
    Create(ctx context.Context, obj *Bucket) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Bucket, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Bucket, new *Bucket, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Bucket, new *Bucket, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// BucketProvider is a dynamic gRPC-based plugin for managing Bucket resources.
type BucketProvider struct {
    ops BucketProviderOps
}

// NewBucketProvider allocates a resource provider that delegates to a ops instance.
func NewBucketProvider(ops BucketProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &BucketProvider{ops: ops}
}

func (p *BucketProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(BucketToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Bucket", "name", failure))
        }
    }
    if !unks["bucketName"] {
        if failure := p.ops.Check(ctx, obj, "bucketName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Bucket", "bucketName", failure))
        }
    }
    if !unks["accessControl"] {
        if failure := p.ops.Check(ctx, obj, "accessControl"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Bucket", "accessControl", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *BucketProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(BucketToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Bucket_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *BucketProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(BucketToken))
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

func (p *BucketProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(BucketToken))
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

func (p *BucketProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(BucketToken))
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
        if diff.Changed("bucketName") {
            replaces = append(replaces, "bucketName")
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

func (p *BucketProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(BucketToken))
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

func (p *BucketProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(BucketToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *BucketProvider) Unmarshal(
    v *pbstruct.Struct) (*Bucket, resource.PropertyMap, error) {
    var obj Bucket
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Bucket structure(s) */

// Bucket is a marshalable representation of its corresponding IDL type.
type Bucket struct {
    Name *string `lumi:"name,optional"`
    BucketName *string `lumi:"bucketName,optional"`
    AccessControl *CannedACL `lumi:"accessControl,optional"`
}

// Bucket's properties have constants to make dealing with diffs and property bags easier.
const (
    Bucket_Name = "name"
    Bucket_BucketName = "bucketName"
    Bucket_AccessControl = "accessControl"
)


