package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderGet get order list
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
func AliexpressSolutionOrderGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderGetAPIRequest, resp *aesolution.AliexpressSolutionOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
