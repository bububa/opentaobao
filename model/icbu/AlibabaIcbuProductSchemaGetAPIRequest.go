package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaGetAPIRequest （新）ICBU商品发布schema接口 API请求
// alibaba.icbu.product.schema.get
//
// 获取ICBU商品发布的页面规则和填写字段，适用于新发商品
type AlibabaIcbuProductSchemaGetAPIRequest struct {
	model.Params
	// 商品规则渲染请求
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaGetRequest 初始化AlibabaIcbuProductSchemaGetAPIRequest对象
func NewAlibabaIcbuProductSchemaGetRequest() *AlibabaIcbuProductSchemaGetAPIRequest {
	return &AlibabaIcbuProductSchemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaGetAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaGetAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaGetAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaGetRequest()
	},
}

// GetAlibabaIcbuProductSchemaGetRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaGetAPIRequest
func GetAlibabaIcbuProductSchemaGetAPIRequest() *AlibabaIcbuProductSchemaGetAPIRequest {
	return poolAlibabaIcbuProductSchemaGetAPIRequest.Get().(*AlibabaIcbuProductSchemaGetAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaGetAPIRequest 将 AlibabaIcbuProductSchemaGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaGetAPIRequest(v *AlibabaIcbuProductSchemaGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaGetAPIRequest.Put(v)
}
