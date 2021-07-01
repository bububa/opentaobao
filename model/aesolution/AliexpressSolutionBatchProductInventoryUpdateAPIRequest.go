package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionBatchProductInventoryUpdateAPIRequest
aliexpress.solution.batch.product.inventory.update API请求
aliexpress.solution.batch.product.inventory.update

batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product. */
type AliexpressSolutionBatchProductInventoryUpdateAPIRequest struct {
	model.Params
	// The product list, in which the inventory needs to be updated. Maximum 20 products.
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// New
