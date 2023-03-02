// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeDirectConnectGateways,DescribeDirectConnectGatewayAssociations,DescribeDirectConnectGatewayAssociationProposals -ContextOnly"; DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
)

func describeGatewayAssociationProposalsPages(ctx context.Context, conn directconnectiface.DirectConnectAPI, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeGatewayAssociationsPages(ctx context.Context, conn directconnectiface.DirectConnectAPI, input *directconnect.DescribeDirectConnectGatewayAssociationsInput, fn func(*directconnect.DescribeDirectConnectGatewayAssociationsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewayAssociationsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeGatewaysPages(ctx context.Context, conn directconnectiface.DirectConnectAPI, input *directconnect.DescribeDirectConnectGatewaysInput, fn func(*directconnect.DescribeDirectConnectGatewaysOutput, bool) bool) error {
	for {
		output, err := conn.DescribeDirectConnectGatewaysWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
