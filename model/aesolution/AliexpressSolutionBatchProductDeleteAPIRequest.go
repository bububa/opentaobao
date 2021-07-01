package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionBatchProductDeleteAPIRequest
aliexpress.solution.batch.product.delete API请求
aliexpress.solution.batch.product.delete

Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious. */
type AliexpressSolutionBatchProductDeleteAPIRequest struct {
	model.Params
	// maximum 100
	_productIds []int64
}

// New
