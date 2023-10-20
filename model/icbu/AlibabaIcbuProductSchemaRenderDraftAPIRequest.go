package icbu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaRenderDraftAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.render.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaRenderDraftAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaRenderDraftAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaRenderDraftAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaRenderDraftRequest()
	},
}

// GetAlibabaIcbuProductSchemaRenderDraftRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaRenderDraftAPIRequest
func GetAlibabaIcbuProductSchemaRenderDraftAPIRequest() *AlibabaIcbuProductSchemaRenderDraftAPIRequest {
	return poolAlibabaIcbuProductSchemaRenderDraftAPIRequest.Get().(*AlibabaIcbuProductSchemaRenderDraftAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaRenderDraftAPIRequest 将 AlibabaIcbuProductSchemaRenderDraftAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaRenderDraftAPIRequest(v *AlibabaIcbuProductSchemaRenderDraftAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaRenderDraftAPIRequest.Put(v)
}
