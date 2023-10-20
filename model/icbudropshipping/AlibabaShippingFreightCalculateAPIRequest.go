package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShippingFreightCalculateAPIRequest 阿里巴巴商品运费计算查询接口 API请求
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
type AlibabaShippingFreightCalculateAPIRequest struct {
	model.Params
	// {}
	_paramFreightTemplateRequest *FreightTemplateRequest
}

// NewAlibabaShippingFreightCalculateRequest 初始化AlibabaShippingFreightCalculateAPIRequest对象
func NewAlibabaShippingFreightCalculateRequest() *AlibabaShippingFreightCalculateAPIRequest {
	return &AlibabaShippingFreightCalculateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaShippingFreightCalculateAPIRequest) Reset() {
	r._paramFreightTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShippingFreightCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.shipping.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShippingFreightCalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaShippingFreightCalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFreightTemplateRequest is ParamFreightTemplateRequest Setter
// {}
func (r *AlibabaShippingFreightCalculateAPIRequest) SetParamFreightTemplateRequest(_paramFreightTemplateRequest *FreightTemplateRequest) error {
	r._paramFreightTemplateRequest = _paramFreightTemplateRequest
	r.Set("param_freight_template_request", _paramFreightTemplateRequest)
	return nil
}

// GetParamFreightTemplateRequest ParamFreightTemplateRequest Getter
func (r AlibabaShippingFreightCalculateAPIRequest) GetParamFreightTemplateRequest() *FreightTemplateRequest {
	return r._paramFreightTemplateRequest
}

var poolAlibabaShippingFreightCalculateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaShippingFreightCalculateRequest()
	},
}

// GetAlibabaShippingFreightCalculateRequest 从 sync.Pool 获取 AlibabaShippingFreightCalculateAPIRequest
func GetAlibabaShippingFreightCalculateAPIRequest() *AlibabaShippingFreightCalculateAPIRequest {
	return poolAlibabaShippingFreightCalculateAPIRequest.Get().(*AlibabaShippingFreightCalculateAPIRequest)
}

// ReleaseAlibabaShippingFreightCalculateAPIRequest 将 AlibabaShippingFreightCalculateAPIRequest 放入 sync.Pool
func ReleaseAlibabaShippingFreightCalculateAPIRequest(v *AlibabaShippingFreightCalculateAPIRequest) {
	v.Reset()
	poolAlibabaShippingFreightCalculateAPIRequest.Put(v)
}
