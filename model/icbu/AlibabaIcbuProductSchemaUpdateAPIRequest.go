package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaUpdateAPIRequest （新）商品发布增量更新接口 API请求
// alibaba.icbu.product.schema.update
//
// 商品更新接口，方式为增量更新，只更新传入的字段
type AlibabaIcbuProductSchemaUpdateAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaUpdateRequest 初始化AlibabaIcbuProductSchemaUpdateAPIRequest对象
func NewAlibabaIcbuProductSchemaUpdateRequest() *AlibabaIcbuProductSchemaUpdateAPIRequest {
	return &AlibabaIcbuProductSchemaUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaUpdateAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaUpdateAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaUpdateAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaUpdateRequest()
	},
}

// GetAlibabaIcbuProductSchemaUpdateRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaUpdateAPIRequest
func GetAlibabaIcbuProductSchemaUpdateAPIRequest() *AlibabaIcbuProductSchemaUpdateAPIRequest {
	return poolAlibabaIcbuProductSchemaUpdateAPIRequest.Get().(*AlibabaIcbuProductSchemaUpdateAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaUpdateAPIRequest 将 AlibabaIcbuProductSchemaUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaUpdateAPIRequest(v *AlibabaIcbuProductSchemaUpdateAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaUpdateAPIRequest.Put(v)
}
