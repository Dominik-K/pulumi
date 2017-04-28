// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for VPC resource provider */

// VPCToken is the type token corresponding to the VPC package type.
const VPCToken = tokens.Type("aws:ec2/vpc:VPC")

// VPCProviderOps is a pluggable interface for VPC-related management functionality.
type VPCProviderOps interface {
    Check(ctx context.Context, obj *VPC) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *VPC) (string, error)
    Get(ctx context.Context, id string) (*VPC, error)
    InspectChange(ctx context.Context,
        id string, old *VPC, new *VPC, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id string, old *VPC, new *VPC, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id string) error
}

// VPCProvider is a dynamic gRPC-based plugin for managing VPC resources.
type VPCProvider struct {
    ops VPCProviderOps
}

// NewVPCProvider allocates a resource provider that delegates to a ops instance.
func NewVPCProvider(ops VPCProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &VPCProvider{ops: ops}
}

func (p *VPCProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *VPCProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *VPCProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   id,
    }, nil
}

func (p *VPCProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    id := req.GetId()
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *VPCProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    if diff.Changed("cidrBlock") {
        replaces = append(replaces, "cidrBlock")
    }
    if diff.Changed("instanceTenancy") {
        replaces = append(replaces, "instanceTenancy")
    }
    more, err := p.ops.InspectChange(ctx, req.GetId(), old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *VPCProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, req.GetId(), old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *VPCProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCToken))
    if err := p.ops.Delete(ctx, req.GetId()); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *VPCProvider) Unmarshal(
    v *pbstruct.Struct) (*VPC, resource.PropertyMap, mapper.DecodeError) {
    var obj VPC
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable VPC structure(s) */

// VPC is a marshalable representation of its corresponding IDL type.
type VPC struct {
    Name string `json:"name"`
    CIDRBlock string `json:"cidrBlock"`
    InstanceTenancy *InstanceTenancy `json:"instanceTenancy,omitempty"`
    EnableDNSSupport *bool `json:"enableDnsSupport,omitempty"`
    EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty"`
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

