package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionBatchProductPriceUpdateAPIRequest
aliexpress.solution.batch.product.price.update API请求
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers */
type AliexpressSolutionBatchProductPriceUpdateAPIRequest struct {
	model.Params
	// The product list, in which the price needs to be updated. Maximum length:20
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// New
