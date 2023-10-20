package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductschemarenderdraftAPIRequest （新）渲染草稿商品数据 API请求
// alibaba.icbu.product.schema.render.draft
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
type AlibabaicbuproductschemarenderdraftAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaicbuproductschemarenderdraftRequest 初始化AlibabaicbuproductschemarenderdraftAPIRequest对象
func NewAlibabaicbuproductschemarenderdraftRequest() *AlibabaicbuproductschemarenderdraftAPIRequest {
	return &AlibabaicbuproductschemarenderdraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductschemarenderdraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.render.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductschemarenderdraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductschemarenderdraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaicbuproductschemarenderdraftAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaicbuproductschemarenderdraftAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
