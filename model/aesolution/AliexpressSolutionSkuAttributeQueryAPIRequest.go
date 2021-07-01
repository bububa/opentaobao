package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionSkuAttributeQueryAPIRequest
Query the sku attribute information belonged to a specific category API请求
aliexpress.solution.sku.attribute.query

Query the sku attribute information belonged to a specific category, customized for oversea merchants. */
type AliexpressSolutionSkuAttributeQueryAPIRequest struct {
	model.Params
	// input parameters
	_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest
}

// New
