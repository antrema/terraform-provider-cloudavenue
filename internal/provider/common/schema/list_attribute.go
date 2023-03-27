package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ Attribute = ListAttribute{}

type ListAttribute struct {
	Common     *schemaR.ListAttribute
	Resource   *schemaR.ListAttribute
	DataSource *schemaD.ListAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s ListAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s ListAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

func (s ListAttribute) GetResource(_ context.Context) schemaR.Attribute {
	var a schemaR.ListAttribute

	if s.Common != nil {
		a = schemaR.ListAttribute{
			Required:            s.Common.Required,
			Optional:            s.Common.Optional,
			Computed:            s.Common.Computed,
			MarkdownDescription: s.Common.MarkdownDescription,
			Description:         s.Common.Description,
			DeprecationMessage:  s.Common.DeprecationMessage,
			Validators:          s.Common.Validators,
			PlanModifiers:       s.Common.PlanModifiers,
			Default:             s.Common.Default,
			ElementType:         s.Common.ElementType,
		}
	}

	if s.Resource != nil {
		if s.Resource.Required {
			a.Required = true
		}

		if s.Resource.Optional {
			a.Optional = true
		}

		if s.Resource.Computed {
			a.Computed = true
		}

		if s.Resource.MarkdownDescription != "" {
			a.MarkdownDescription += s.Resource.MarkdownDescription
		}

		if s.Resource.Description != "" {
			a.Description += s.Resource.Description
		}

		if s.Resource.DeprecationMessage != "" {
			a.DeprecationMessage += s.Resource.DeprecationMessage
		}

		if len(s.Resource.Validators) > 0 {
			a.Validators = append(a.Validators, s.Resource.Validators...)
		}

		if len(s.Resource.PlanModifiers) > 0 {
			a.PlanModifiers = append(a.PlanModifiers, s.Resource.PlanModifiers...)
		}

		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
		}

		if s.Resource.ElementType != nil {
			a.ElementType = s.Resource.ElementType
		}
	}

	return a
}

func (s ListAttribute) GetDataSource(_ context.Context) schemaD.Attribute {
	var a schemaD.ListAttribute

	if s.Common != nil {
		a = schemaD.ListAttribute{
			Required:            s.Common.Required,
			Optional:            s.Common.Optional,
			Computed:            s.Common.Computed,
			MarkdownDescription: s.Common.MarkdownDescription,
			Description:         s.Common.Description,
			DeprecationMessage:  s.Common.DeprecationMessage,
			Validators:          s.Common.Validators,
			ElementType:         s.Common.ElementType,
		}
	}

	if s.DataSource != nil {
		if s.DataSource.Required {
			a.Required = true
		}

		if s.DataSource.Optional {
			a.Optional = true
		}

		if s.DataSource.Computed {
			a.Computed = true
		}

		if s.DataSource.MarkdownDescription != "" {
			a.MarkdownDescription += s.DataSource.MarkdownDescription
		}

		if s.DataSource.Description != "" {
			a.Description += s.DataSource.Description
		}

		if s.DataSource.DeprecationMessage != "" {
			a.DeprecationMessage += s.DataSource.DeprecationMessage
		}

		if len(s.DataSource.Validators) > 0 {
			a.Validators = append(a.Validators, s.DataSource.Validators...)
		}

		if s.DataSource.ElementType != nil {
			a.ElementType = s.DataSource.ElementType
		}
	}

	return a
}