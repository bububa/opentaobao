package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemaupdateAPIRequest （新）商品发布增量更新接口 API请求
// alibaba.icbu.product.schema.update
//
// 商品更新接口，方式为增量更新，只更新传入的字段
type AlibabaicbuproductschemaupdateAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaicbuproductschemaupdateRequest 初始化AlibabaicbuproductschemaupdateAPIRequest对象
func NewAlibabaicbuproductschemaupdateRequest() *AlibabaicbuproductschemaupdateAPIRequest {
	return &AlibabaicbuproductschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductschemaupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaicbuproductschemaupdateAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaicbuproductschemaupdateAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
