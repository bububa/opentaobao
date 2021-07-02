package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaRenderDraftAPIRequest （新）渲染草稿商品数据 API请求
// alibaba.icbu.product.schema.render.draft
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
type AlibabaIcbuProductSchemaRenderDraftAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaRenderDraftRequest 初始化AlibabaIcbuProductSchemaRenderDraftAPIRequest对象
func NewAlibabaIcbuProductSchemaRenderDraftRequest() *AlibabaIcbuProductSchemaRenderDraftAPIRequest {
	return &AlibabaIcbuProductSchemaRenderDraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.render.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaRenderDraftAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// Get ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}
