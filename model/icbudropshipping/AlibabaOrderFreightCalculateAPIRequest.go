package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaorderfreightcalculateAPIRequest 阿里巴巴下单场景运费方案计算 API请求
// alibaba.order.freight.calculate
//
// icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
// alibaba Create order scenario freight calculation
type AlibabaorderfreightcalculateAPIRequest struct {
	model.Params
	// {}
	_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest
}

// NewAlibabaorderfreightcalculateRequest 初始化AlibabaorderfreightcalculateAPIRequest对象
func NewAlibabaorderfreightcalculateRequest() *AlibabaorderfreightcalculateAPIRequest {
	return &AlibabaorderfreightcalculateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaorderfreightcalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.order.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaorderfreightcalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaorderfreightcalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMultiFreightTemplateRequest is ParamMultiFreightTemplateRequest Setter
// {}
func (r *AlibabaorderfreightcalculateAPIRequest) SetParamMultiFreightTemplateRequest(_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest) error {
	r._paramMultiFreightTemplateRequest = _paramMultiFreightTemplateRequest
	r.Set("param_multi_freight_template_request", _paramMultiFreightTemplateRequest)
	return nil
}

// GetParamMultiFreightTemplateRequest ParamMultiFreightTemplateRequest Getter
func (r AlibabaorderfreightcalculateAPIRequest) GetParamMultiFreightTemplateRequest() *MultiFreightTemplateRequest {
	return r._paramMultiFreightTemplateRequest
}
