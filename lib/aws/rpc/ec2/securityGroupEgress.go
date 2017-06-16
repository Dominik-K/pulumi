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

/* RPC stubs for SecurityGroupEgress resource provider */

// SecurityGroupEgressToken is the type token corresponding to the SecurityGroupEgress package type.
const SecurityGroupEgressToken = tokens.Type("aws:ec2/securityGroupEgress:SecurityGroupEgress")

// SecurityGroupEgressProviderOps is a pluggable interface for SecurityGroupEgress-related management functionality.
type SecurityGroupEgressProviderOps interface {
    Check(ctx context.Context, obj *SecurityGroupEgress, property string) error
    Create(ctx context.Context, obj *SecurityGroupEgress) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*SecurityGroupEgress, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *SecurityGroupEgress, new *SecurityGroupEgress, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *SecurityGroupEgress, new *SecurityGroupEgress, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SecurityGroupEgressProvider is a dynamic gRPC-based plugin for managing SecurityGroupEgress resources.
type SecurityGroupEgressProvider struct {
    ops SecurityGroupEgressProviderOps
}

// NewSecurityGroupEgressProvider allocates a resource provider that delegates to a ops instance.
func NewSecurityGroupEgressProvider(ops SecurityGroupEgressProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SecurityGroupEgressProvider{ops: ops}
}

func (p *SecurityGroupEgressProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "name", failure))
        }
    }
    if !unks["fromPort"] {
        if failure := p.ops.Check(ctx, obj, "fromPort"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "fromPort", failure))
        }
    }
    if !unks["group"] {
        if failure := p.ops.Check(ctx, obj, "group"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "group", failure))
        }
    }
    if !unks["ipProtocol"] {
        if failure := p.ops.Check(ctx, obj, "ipProtocol"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "ipProtocol", failure))
        }
    }
    if !unks["toPort"] {
        if failure := p.ops.Check(ctx, obj, "toPort"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "toPort", failure))
        }
    }
    if !unks["cidrIp"] {
        if failure := p.ops.Check(ctx, obj, "cidrIp"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "cidrIp", failure))
        }
    }
    if !unks["cidrIpv6"] {
        if failure := p.ops.Check(ctx, obj, "cidrIpv6"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "cidrIpv6", failure))
        }
    }
    if !unks["destinationPrefixListId"] {
        if failure := p.ops.Check(ctx, obj, "destinationPrefixListId"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "destinationPrefixListId", failure))
        }
    }
    if !unks["destinationSecurityGroup"] {
        if failure := p.ops.Check(ctx, obj, "destinationSecurityGroup"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("SecurityGroupEgress", "destinationSecurityGroup", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *SecurityGroupEgressProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[SecurityGroupEgress_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *SecurityGroupEgressProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
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

func (p *SecurityGroupEgressProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
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

func (p *SecurityGroupEgressProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
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
        if diff.Changed("fromPort") {
            replaces = append(replaces, "fromPort")
        }
        if diff.Changed("group") {
            replaces = append(replaces, "group")
        }
        if diff.Changed("ipProtocol") {
            replaces = append(replaces, "ipProtocol")
        }
        if diff.Changed("toPort") {
            replaces = append(replaces, "toPort")
        }
        if diff.Changed("cidrIp") {
            replaces = append(replaces, "cidrIp")
        }
        if diff.Changed("cidrIpv6") {
            replaces = append(replaces, "cidrIpv6")
        }
        if diff.Changed("destinationPrefixListId") {
            replaces = append(replaces, "destinationPrefixListId")
        }
        if diff.Changed("destinationSecurityGroup") {
            replaces = append(replaces, "destinationSecurityGroup")
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

func (p *SecurityGroupEgressProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
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

func (p *SecurityGroupEgressProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupEgressToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupEgressProvider) Unmarshal(
    v *pbstruct.Struct) (*SecurityGroupEgress, resource.PropertyMap, error) {
    var obj SecurityGroupEgress
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable SecurityGroupEgress structure(s) */

// SecurityGroupEgress is a marshalable representation of its corresponding IDL type.
type SecurityGroupEgress struct {
    Name *string `lumi:"name,optional"`
    FromPort float64 `lumi:"fromPort"`
    Group resource.ID `lumi:"group"`
    IPProtocol string `lumi:"ipProtocol"`
    ToPort float64 `lumi:"toPort"`
    CIDRIP *string `lumi:"cidrIp,optional"`
    CIDRIPv6 *string `lumi:"cidrIpv6,optional"`
    DestinationPrefixListId *string `lumi:"destinationPrefixListId,optional"`
    DestinationSecurityGroup *resource.ID `lumi:"destinationSecurityGroup,optional"`
}

// SecurityGroupEgress's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroupEgress_Name = "name"
    SecurityGroupEgress_FromPort = "fromPort"
    SecurityGroupEgress_Group = "group"
    SecurityGroupEgress_IPProtocol = "ipProtocol"
    SecurityGroupEgress_ToPort = "toPort"
    SecurityGroupEgress_CIDRIP = "cidrIp"
    SecurityGroupEgress_CIDRIPv6 = "cidrIpv6"
    SecurityGroupEgress_DestinationPrefixListId = "destinationPrefixListId"
    SecurityGroupEgress_DestinationSecurityGroup = "destinationSecurityGroup"
)


