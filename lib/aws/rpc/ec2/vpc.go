// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

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

/* RPC stubs for VPC resource provider */

// VPCToken is the type token corresponding to the VPC package type.
const VPCToken = tokens.Type("aws:ec2/vpc:VPC")

// VPCProviderOps is a pluggable interface for VPC-related management functionality.
type VPCProviderOps interface {
    Check(ctx context.Context, obj *VPC, property string) error
    Create(ctx context.Context, obj *VPC) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*VPC, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *VPC, new *VPC, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *VPC, new *VPC, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// VPCProvider is a dynamic gRPC-based plugin for managing VPC resources.
type VPCProvider struct {
    ops VPCProviderOps
}

// NewVPCProvider allocates a resource provider that delegates to a ops instance.
func NewVPCProvider(ops VPCProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &VPCProvider{ops: ops}
}

func (p *VPCProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("VPC", "name", failure))
        }
    }
    if !unks["cidrBlock"] {
        if failure := p.ops.Check(ctx, obj, "cidrBlock"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("VPC", "cidrBlock", failure))
        }
    }
    if !unks["instanceTenancy"] {
        if failure := p.ops.Check(ctx, obj, "instanceTenancy"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("VPC", "instanceTenancy", failure))
        }
    }
    if !unks["enableDnsSupport"] {
        if failure := p.ops.Check(ctx, obj, "enableDnsSupport"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("VPC", "enableDnsSupport", failure))
        }
    }
    if !unks["enableDnsHostnames"] {
        if failure := p.ops.Check(ctx, obj, "enableDnsHostnames"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("VPC", "enableDnsHostnames", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *VPCProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[VPC_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *VPCProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
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

func (p *VPCProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
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

func (p *VPCProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
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
        if diff.Changed("cidrBlock") {
            replaces = append(replaces, "cidrBlock")
        }
        if diff.Changed("instanceTenancy") {
            replaces = append(replaces, "instanceTenancy")
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

func (p *VPCProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCToken))
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

func (p *VPCProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *VPCProvider) Unmarshal(
    v *pbstruct.Struct) (*VPC, resource.PropertyMap, error) {
    var obj VPC
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable VPC structure(s) */

// VPC is a marshalable representation of its corresponding IDL type.
type VPC struct {
    Name *string `lumi:"name,optional"`
    CIDRBlock string `lumi:"cidrBlock"`
    InstanceTenancy *InstanceTenancy `lumi:"instanceTenancy,optional"`
    EnableDNSSupport *bool `lumi:"enableDnsSupport,optional"`
    EnableDNSHostnames *bool `lumi:"enableDnsHostnames,optional"`
}

// VPC's properties have constants to make dealing with diffs and property bags easier.
const (
    VPC_Name = "name"
    VPC_CIDRBlock = "cidrBlock"
    VPC_InstanceTenancy = "instanceTenancy"
    VPC_EnableDNSSupport = "enableDnsSupport"
    VPC_EnableDNSHostnames = "enableDnsHostnames"
)

/* Typedefs */

type (
    InstanceTenancy string
)

/* Constants */

const (
    DedicatedTenancy InstanceTenancy = "dedicated"
    DefaultTenancy InstanceTenancy = "default"
    HostTenancy InstanceTenancy = "host"
)


