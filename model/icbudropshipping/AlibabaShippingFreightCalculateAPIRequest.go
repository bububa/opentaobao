package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabashippingfreightcalculateAPIRequest 阿里巴巴商品运费计算查询接口 API请求
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
type AlibabashippingfreightcalculateAPIRequest struct {
	model.Params
	// {}
	_paramFreightTemplateRequest *FreightTemplateRequest
}

// NewAlibabashippingfreightcalculateRequest 初始化AlibabashippingfreightcalculateAPIRequest对象
func NewAlibabashippingfreightcalculateRequest() *AlibabashippingfreightcalculateAPIRequest {
	return &AlibabashippingfreightcalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabashippingfreightcalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.shipping.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabashippingfreightcalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabashippingfreightcalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamFreightTemplateRequest is ParamFreightTemplateRequest Setter
// {}
func (r *AlibabashippingfreightcalculateAPIRequest) SetParamFreightTemplateRequest(_paramFreightTemplateRequest *FreightTemplateRequest) error {
	r._paramFreightTemplateRequest = _paramFreightTemplateRequest
	r.Set("param_freight_template_request", _paramFreightTemplateRequest)
	return nil
}

// GetParamFreightTemplateRequest ParamFreightTemplateRequest Getter
func (r AlibabashippingfreightcalculateAPIRequest) GetParamFreightTemplateRequest() *FreightTemplateRequest {
	return r._paramFreightTemplateRequest
}
