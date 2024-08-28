package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderFulfill fulfill order
// aliexpress.solution.order.fulfill
//
// fulfill order for seller
func AliexpressSolutionOrderFulfill(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderFulfillAPIRequest, resp *aesolution.AliexpressSolutionOrderFulfillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
