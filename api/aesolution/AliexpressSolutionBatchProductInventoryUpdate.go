package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionBatchProductInventoryUpdate aliexpress.solution.batch.product.inventory.update
// aliexpress.solution.batch.product.inventory.update
//
// batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
func AliexpressSolutionBatchProductInventoryUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionBatchProductInventoryUpdateAPIRequest, session string) (*aesolution.AliexpressSolutionBatchProductInventoryUpdateAPIResponse, error) {
	var resp aesolution.AliexpressSolutionBatchProductInventoryUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
