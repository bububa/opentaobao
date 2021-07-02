package icbudropshipping

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShippingFreightCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.shipping.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShippingFreightCalculateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
