package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionProductInfoGetAPIRequest
Get Single Product Info API请求
aliexpress.solution.product.info.get

Get Single Product Info */
type AliexpressSolutionProductInfoGetAPIRequest struct {
	model.Params
	// product ID
	_productId int64
}

// New
