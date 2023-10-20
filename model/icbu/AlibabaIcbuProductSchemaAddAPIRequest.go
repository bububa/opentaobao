package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaAddAPIRequest （新）商品发布新接口 API请求
// alibaba.icbu.product.schema.add
//
// 提供发布ICBU商品的入口
type AlibabaIcbuProductSchemaAddAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaAddRequest 初始化AlibabaIcbuProductSchemaAddAPIRequest对象
func NewAlibabaIcbuProductSchemaAddRequest() *AlibabaIcbuProductSchemaAddAPIRequest {
	return &AlibabaIcbuProductSchemaAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaAddAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaAddAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaAddAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaAddRequest()
	},
}

// GetAlibabaIcbuProductSchemaAddRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaAddAPIRequest
func GetAlibabaIcbuProductSchemaAddAPIRequest() *AlibabaIcbuProductSchemaAddAPIRequest {
	return poolAlibabaIcbuProductSchemaAddAPIRequest.Get().(*AlibabaIcbuProductSchemaAddAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaAddAPIRequest 将 AlibabaIcbuProductSchemaAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaAddAPIRequest(v *AlibabaIcbuProductSchemaAddAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaAddAPIRequest.Put(v)
}
