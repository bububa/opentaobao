package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionProductListGetAPIRequest
Get product list API请求
aliexpress.solution.product.list.get

Get product list */
type AliexpressSolutionProductListGetAPIRequest struct {
	model.Params
	// request parameters to query
	_aeopAEProductListQuery *ItemListQuery
}

// New
