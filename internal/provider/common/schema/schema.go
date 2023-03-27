package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type Schema struct {
	Common     SchemaDetails
	Resource   SchemaDetails
	DataSource SchemaDetails
	Attributes Attributes
}

type SchemaDetails struct {
	MarkdownDescription string
	DeprecationMessage  string
}

func (s Schema) GetResource(ctx context.Context) schemaR.Schema {
	if s.Resource.MarkdownDescription != "" {
		s.Common.MarkdownDescription += s.Resource.MarkdownDescription
	}

	if s.Resource.DeprecationMessage != "" {
		s.Common.DeprecationMessage += s.Resource.DeprecationMessage
	}

	return schemaR.Schema{
		MarkdownDescription: s.Common.MarkdownDescription,
		DeprecationMessage:  s.Common.DeprecationMessage,
		Attributes:          s.Attributes.process(ctx, resource).(map[string]schemaR.Attribute),
	}
}

func (s Schema) GetDataSource(ctx context.Context) schemaD.Schema {
	if s.DataSource.MarkdownDescription != "" {
		s.Common.MarkdownDescription += s.DataSource.MarkdownDescription
	}

	if s.DataSource.DeprecationMessage != "" {
		s.Common.DeprecationMessage += s.DataSource.DeprecationMessage
	}

	return schemaD.Schema{
		MarkdownDescription: s.Common.MarkdownDescription,
		DeprecationMessage:  s.Common.DeprecationMessage,
		Attributes:          s.Attributes.process(ctx, dataSource).(map[string]schemaD.Attribute),
	}
}