package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemarenderAPIRequest （新）获取商品信息 API请求
// alibaba.icbu.product.schema.render
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
type AlibabaicbuproductschemarenderAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaicbuproductschemarenderRequest 初始化AlibabaicbuproductschemarenderAPIRequest对象
func NewAlibabaicbuproductschemarenderRequest() *AlibabaicbuproductschemarenderAPIRequest {
	return &AlibabaicbuproductschemarenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductschemarenderAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductschemarenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductschemarenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaicbuproductschemarenderAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaicbuproductschemarenderAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
