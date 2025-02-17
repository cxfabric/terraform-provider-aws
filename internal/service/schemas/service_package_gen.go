// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package schemas

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	schemas_sdkv1 "github.com/aws/aws-sdk-go/service/schemas"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDiscoverer,
			TypeName: "aws_schemas_discoverer",
			Name:     "Discoverer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRegistry,
			TypeName: "aws_schemas_registry",
			Name:     "Registry",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRegistryPolicy,
			TypeName: "aws_schemas_registry_policy",
		},
		{
			Factory:  ResourceSchema,
			TypeName: "aws_schemas_schema",
			Name:     "Schema",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Schemas
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*schemas_sdkv1.Schemas, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return schemas_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
