package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionBatchProductInventoryUpdate aliexpress.solution.batch.product.inventory.update
// aliexpress.solution.batch.product.inventory.update
//
// batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
func AliexpressSolutionBatchProductInventoryUpdate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionBatchProductInventoryUpdateAPIRequest, resp *aesolution.AliexpressSolutionBatchProductInventoryUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
