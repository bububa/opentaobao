package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemaadddraftAPIRequest （新）ICBU商品发布草稿 API请求
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
type AlibabaicbuproductschemaadddraftAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaicbuproductschemaadddraftRequest 初始化AlibabaicbuproductschemaadddraftAPIRequest对象
func NewAlibabaicbuproductschemaadddraftRequest() *AlibabaicbuproductschemaadddraftAPIRequest {
	return &AlibabaicbuproductschemaadddraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductschemaadddraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.add.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductschemaadddraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductschemaadddraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaicbuproductschemaadddraftAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaicbuproductschemaadddraftAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
