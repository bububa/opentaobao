package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaRenderAPIRequest （新）获取商品信息 API请求
// alibaba.icbu.product.schema.render
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
type AlibabaIcbuProductSchemaRenderAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaRenderRequest 初始化AlibabaIcbuProductSchemaRenderAPIRequest对象
func NewAlibabaIcbuProductSchemaRenderRequest() *AlibabaIcbuProductSchemaRenderAPIRequest {
	return &AlibabaIcbuProductSchemaRenderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaRenderAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaRenderAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaRenderAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaRenderAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaRenderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaRenderRequest()
	},
}

// GetAlibabaIcbuProductSchemaRenderRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaRenderAPIRequest
func GetAlibabaIcbuProductSchemaRenderAPIRequest() *AlibabaIcbuProductSchemaRenderAPIRequest {
	return poolAlibabaIcbuProductSchemaRenderAPIRequest.Get().(*AlibabaIcbuProductSchemaRenderAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaRenderAPIRequest 将 AlibabaIcbuProductSchemaRenderAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaRenderAPIRequest(v *AlibabaIcbuProductSchemaRenderAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaRenderAPIRequest.Put(v)
}
