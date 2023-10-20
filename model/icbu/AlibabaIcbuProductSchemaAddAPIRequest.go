package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemaaddAPIRequest （新）商品发布新接口 API请求
// alibaba.icbu.product.schema.add
//
// 提供发布ICBU商品的入口
type AlibabaicbuproductschemaaddAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaicbuproductschemaaddRequest 初始化AlibabaicbuproductschemaaddAPIRequest对象
func NewAlibabaicbuproductschemaaddRequest() *AlibabaicbuproductschemaaddAPIRequest {
	return &AlibabaicbuproductschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductschemaaddAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaicbuproductschemaaddAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaicbuproductschemaaddAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
