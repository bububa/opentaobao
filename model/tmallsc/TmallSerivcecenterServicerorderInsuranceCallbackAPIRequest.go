package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallserivcecenterservicerorderinsurancecallbackAPIRequest 服务商回传保单信息 API请求
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
type TmallserivcecenterservicerorderinsurancecallbackAPIRequest struct {
	model.Params
	// 服务单投保信息回传入参
	_serviceInsuranceCallbackRequest *ServiceInsuranceCallbackRequest
}

// NewTmallserivcecenterservicerorderinsurancecallbackRequest 初始化TmallserivcecenterservicerorderinsurancecallbackAPIRequest对象
func NewTmallserivcecenterservicerorderinsurancecallbackRequest() *TmallserivcecenterservicerorderinsurancecallbackAPIRequest {
	return &TmallserivcecenterservicerorderinsurancecallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallserivcecenterservicerorderinsurancecallbackAPIRequest) GetApiMethodName() string {
	return "tmall.serivcecenter.servicerorder.insurance.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallserivcecenterservicerorderinsurancecallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallserivcecenterservicerorderinsurancecallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceInsuranceCallbackRequest is ServiceInsuranceCallbackRequest Setter
// 服务单投保信息回传入参
func (r *TmallserivcecenterservicerorderinsurancecallbackAPIRequest) SetServiceInsuranceCallbackRequest(_serviceInsuranceCallbackRequest *ServiceInsuranceCallbackRequest) error {
	r._serviceInsuranceCallbackRequest = _serviceInsuranceCallbackRequest
	r.Set("service_insurance_callback_request", _serviceInsuranceCallbackRequest)
	return nil
}

// GetServiceInsuranceCallbackRequest ServiceInsuranceCallbackRequest Getter
func (r TmallserivcecenterservicerorderinsurancecallbackAPIRequest) GetServiceInsuranceCallbackRequest() *ServiceInsuranceCallbackRequest {
	return r._serviceInsuranceCallbackRequest
}
