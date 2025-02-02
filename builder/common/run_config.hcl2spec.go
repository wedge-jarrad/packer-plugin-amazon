// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAmiFilterOptions struct {
	Filters    map[string]string `mapstructure:"filters" cty:"filters" hcl:"filters"`
	Owners     []string          `mapstructure:"owners" cty:"owners" hcl:"owners"`
	MostRecent *bool             `mapstructure:"most_recent" cty:"most_recent" hcl:"most_recent"`
}

// FlatMapstructure returns a new FlatAmiFilterOptions.
// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AmiFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatAmiFilterOptions)
}

// HCL2Spec returns the hcl spec of a AmiFilterOptions.
// This spec is used by HCL to read the fields of AmiFilterOptions.
// The decoded values from this spec will then be applied to a FlatAmiFilterOptions.
func (*FlatAmiFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":     &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"owners":      &hcldec.AttrSpec{Name: "owners", Type: cty.List(cty.String), Required: false},
		"most_recent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatLicenseConfigurationRequest is an auto-generated flat version of LicenseConfigurationRequest.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatLicenseConfigurationRequest struct {
	LicenseConfigurationArn *string `mapstructure:"license_configuration_arn" cty:"license_configuration_arn" hcl:"license_configuration_arn"`
}

// FlatMapstructure returns a new FlatLicenseConfigurationRequest.
// FlatLicenseConfigurationRequest is an auto-generated flat version of LicenseConfigurationRequest.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*LicenseConfigurationRequest) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatLicenseConfigurationRequest)
}

// HCL2Spec returns the hcl spec of a LicenseConfigurationRequest.
// This spec is used by HCL to read the fields of LicenseConfigurationRequest.
// The decoded values from this spec will then be applied to a FlatLicenseConfigurationRequest.
func (*FlatLicenseConfigurationRequest) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"license_configuration_arn": &hcldec.AttrSpec{Name: "license_configuration_arn", Type: cty.String, Required: false},
	}
	return s
}

// FlatLicenseSpecification is an auto-generated flat version of LicenseSpecification.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatLicenseSpecification struct {
	LicenseConfigurationRequest *FlatLicenseConfigurationRequest `mapstructure:"license_configuration_request" cty:"license_configuration_request" hcl:"license_configuration_request"`
}

// FlatMapstructure returns a new FlatLicenseSpecification.
// FlatLicenseSpecification is an auto-generated flat version of LicenseSpecification.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*LicenseSpecification) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatLicenseSpecification)
}

// HCL2Spec returns the hcl spec of a LicenseSpecification.
// This spec is used by HCL to read the fields of LicenseSpecification.
// The decoded values from this spec will then be applied to a FlatLicenseSpecification.
func (*FlatLicenseSpecification) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"license_configuration_request": &hcldec.BlockSpec{TypeName: "license_configuration_request", Nested: hcldec.ObjectSpec((*FlatLicenseConfigurationRequest)(nil).HCL2Spec())},
	}
	return s
}

// FlatMetadataOptions is an auto-generated flat version of MetadataOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatMetadataOptions struct {
	HttpEndpoint            *string `mapstructure:"http_endpoint" required:"false" cty:"http_endpoint" hcl:"http_endpoint"`
	HttpTokens              *string `mapstructure:"http_tokens" required:"false" cty:"http_tokens" hcl:"http_tokens"`
	HttpPutResponseHopLimit *int64  `mapstructure:"http_put_response_hop_limit" required:"false" cty:"http_put_response_hop_limit" hcl:"http_put_response_hop_limit"`
}

// FlatMapstructure returns a new FlatMetadataOptions.
// FlatMetadataOptions is an auto-generated flat version of MetadataOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*MetadataOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatMetadataOptions)
}

// HCL2Spec returns the hcl spec of a MetadataOptions.
// This spec is used by HCL to read the fields of MetadataOptions.
// The decoded values from this spec will then be applied to a FlatMetadataOptions.
func (*FlatMetadataOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"http_endpoint":               &hcldec.AttrSpec{Name: "http_endpoint", Type: cty.String, Required: false},
		"http_tokens":                 &hcldec.AttrSpec{Name: "http_tokens", Type: cty.String, Required: false},
		"http_put_response_hop_limit": &hcldec.AttrSpec{Name: "http_put_response_hop_limit", Type: cty.Number, Required: false},
	}
	return s
}

// FlatPlacement is an auto-generated flat version of Placement.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatPlacement struct {
	HostResourceGroupArn *string `mapstructure:"host_resource_group_arn" required:"false" cty:"host_resource_group_arn" hcl:"host_resource_group_arn"`
	Tenancy              *string `mapstructure:"tenancy" required:"false" cty:"tenancy" hcl:"tenancy"`
}

// FlatMapstructure returns a new FlatPlacement.
// FlatPlacement is an auto-generated flat version of Placement.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Placement) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatPlacement)
}

// HCL2Spec returns the hcl spec of a Placement.
// This spec is used by HCL to read the fields of Placement.
// The decoded values from this spec will then be applied to a FlatPlacement.
func (*FlatPlacement) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"host_resource_group_arn": &hcldec.AttrSpec{Name: "host_resource_group_arn", Type: cty.String, Required: false},
		"tenancy":                 &hcldec.AttrSpec{Name: "tenancy", Type: cty.String, Required: false},
	}
	return s
}

// FlatPolicyDocument is an auto-generated flat version of PolicyDocument.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatPolicyDocument struct {
	Version   *string         `mapstructure:"Version" required:"false" cty:"Version" hcl:"Version"`
	Statement []FlatStatement `mapstructure:"Statement" required:"false" cty:"Statement" hcl:"Statement"`
}

// FlatMapstructure returns a new FlatPolicyDocument.
// FlatPolicyDocument is an auto-generated flat version of PolicyDocument.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*PolicyDocument) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatPolicyDocument)
}

// HCL2Spec returns the hcl spec of a PolicyDocument.
// This spec is used by HCL to read the fields of PolicyDocument.
// The decoded values from this spec will then be applied to a FlatPolicyDocument.
func (*FlatPolicyDocument) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Version":   &hcldec.AttrSpec{Name: "Version", Type: cty.String, Required: false},
		"Statement": &hcldec.BlockListSpec{TypeName: "Statement", Nested: hcldec.ObjectSpec((*FlatStatement)(nil).HCL2Spec())},
	}
	return s
}

// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSecurityGroupFilterOptions struct {
	Filters map[string]string      `cty:"filters" hcl:"filters"`
	Filter  []config.FlatNameValue `cty:"filter" hcl:"filter"`
}

// FlatMapstructure returns a new FlatSecurityGroupFilterOptions.
// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SecurityGroupFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSecurityGroupFilterOptions)
}

// HCL2Spec returns the hcl spec of a SecurityGroupFilterOptions.
// This spec is used by HCL to read the fields of SecurityGroupFilterOptions.
// The decoded values from this spec will then be applied to a FlatSecurityGroupFilterOptions.
func (*FlatSecurityGroupFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":  &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*config.FlatNameValue)(nil).HCL2Spec())},
	}
	return s
}

// FlatStatement is an auto-generated flat version of Statement.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatStatement struct {
	Effect   *string  `mapstructure:"Effect" required:"false" cty:"Effect" hcl:"Effect"`
	Action   []string `mapstructure:"Action" required:"false" cty:"Action" hcl:"Action"`
	Resource []string `mapstructure:"Resource" required:"false" cty:"Resource" hcl:"Resource"`
}

// FlatMapstructure returns a new FlatStatement.
// FlatStatement is an auto-generated flat version of Statement.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Statement) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatStatement)
}

// HCL2Spec returns the hcl spec of a Statement.
// This spec is used by HCL to read the fields of Statement.
// The decoded values from this spec will then be applied to a FlatStatement.
func (*FlatStatement) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"Effect":   &hcldec.AttrSpec{Name: "Effect", Type: cty.String, Required: false},
		"Action":   &hcldec.AttrSpec{Name: "Action", Type: cty.List(cty.String), Required: false},
		"Resource": &hcldec.AttrSpec{Name: "Resource", Type: cty.List(cty.String), Required: false},
	}
	return s
}

// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSubnetFilterOptions struct {
	Filters  map[string]string      `cty:"filters" hcl:"filters"`
	Filter   []config.FlatNameValue `cty:"filter" hcl:"filter"`
	MostFree *bool                  `mapstructure:"most_free" cty:"most_free" hcl:"most_free"`
	Random   *bool                  `mapstructure:"random" cty:"random" hcl:"random"`
}

// FlatMapstructure returns a new FlatSubnetFilterOptions.
// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SubnetFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSubnetFilterOptions)
}

// HCL2Spec returns the hcl spec of a SubnetFilterOptions.
// This spec is used by HCL to read the fields of SubnetFilterOptions.
// The decoded values from this spec will then be applied to a FlatSubnetFilterOptions.
func (*FlatSubnetFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":   &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":    &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*config.FlatNameValue)(nil).HCL2Spec())},
		"most_free": &hcldec.AttrSpec{Name: "most_free", Type: cty.Bool, Required: false},
		"random":    &hcldec.AttrSpec{Name: "random", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatVpcFilterOptions struct {
	Filters map[string]string      `cty:"filters" hcl:"filters"`
	Filter  []config.FlatNameValue `cty:"filter" hcl:"filter"`
}

// FlatMapstructure returns a new FlatVpcFilterOptions.
// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*VpcFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatVpcFilterOptions)
}

// HCL2Spec returns the hcl spec of a VpcFilterOptions.
// This spec is used by HCL to read the fields of VpcFilterOptions.
// The decoded values from this spec will then be applied to a FlatVpcFilterOptions.
func (*FlatVpcFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":  &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*config.FlatNameValue)(nil).HCL2Spec())},
	}
	return s
}
