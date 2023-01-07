// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package comment

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	Comment  *string `mapstructure:"comment" cty:"comment" hcl:"comment"`
	SendToUi *bool   `mapstructure:"ui" cty:"ui" hcl:"ui"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"comment": &hcldec.AttrSpec{Name: "comment", Type: cty.String, Required: false},
		"ui":      &hcldec.AttrSpec{Name: "ui", Type: cty.Bool, Required: false},
	}
	return s
}