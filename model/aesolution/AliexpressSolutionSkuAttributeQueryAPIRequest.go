package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionskuattributequeryAPIRequest Query the sku attribute information belonged to a specific category API请求
// aliexpress.solution.sku.attribute.query
//
// Query the sku attribute information belonged to a specific category, customized for oversea merchants.
type AliexpresssolutionskuattributequeryAPIRequest struct {
	model.Params
	// input parameters
	_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest
}

// NewAliexpresssolutionskuattributequeryRequest 初始化AliexpresssolutionskuattributequeryAPIRequest对象
func NewAliexpresssolutionskuattributequeryRequest() *AliexpresssolutionskuattributequeryAPIRequest {
	return &AliexpresssolutionskuattributequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionskuattributequeryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.sku.attribute.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionskuattributequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionskuattributequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuerySkuAttributeInfoRequest is QuerySkuAttributeInfoRequest Setter
// input parameters
func (r *AliexpresssolutionskuattributequeryAPIRequest) SetQuerySkuAttributeInfoRequest(_querySkuAttributeInfoRequest *SkuAttributeInfoQueryRequest) error {
	r._querySkuAttributeInfoRequest = _querySkuAttributeInfoRequest
	r.Set("query_sku_attribute_info_request", _querySkuAttributeInfoRequest)
	return nil
}

// GetQuerySkuAttributeInfoRequest QuerySkuAttributeInfoRequest Getter
func (r AliexpresssolutionskuattributequeryAPIRequest) GetQuerySkuAttributeInfoRequest() *SkuAttributeInfoQueryRequest {
	return r._querySkuAttributeInfoRequest
}
