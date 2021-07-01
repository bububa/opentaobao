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

// NewAliexpressSolutionSkuAttributeQueryRequest 初始化AliexpressSolutionSkuAttributeQueryAPIRequest对象
func NewAliexpressSolutionSkuAttributeQueryRequest() *AliexpressSolutionSkuAttributeQueryAPIRequest {
	return &AliexpressSolutionSkuAttributeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.sku.attribute.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QuerySkuAttributeInfoRequest Setter
// input parameters
func (r *AliexpressSolutionSkuAttributeQueryAPIRequest) SetQuerySkuAttributeInfoRequest(_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest) error {
	r._querySkuAttributeInfoRequest = _querySkuAttributeInfoRequest
	r.Set("query_sku_attribute_info_request", _querySkuAttributeInfoRequest)
	return nil
}

// Get QuerySkuAttributeInfoRequest Getter
func (r AliexpressSolutionSkuAttributeQueryAPIRequest) GetQuerySkuAttributeInfoRequest() *SkuAttributeInfoQueryRequest {
	return r._querySkuAttributeInfoRequest
}
