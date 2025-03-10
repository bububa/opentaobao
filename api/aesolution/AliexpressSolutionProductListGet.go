package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionProductListGet Get product list
// aliexpress.solution.product.list.get
//
// Get product list
func AliexpressSolutionProductListGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionProductListGetAPIRequest, resp *aesolution.AliexpressSolutionProductListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
