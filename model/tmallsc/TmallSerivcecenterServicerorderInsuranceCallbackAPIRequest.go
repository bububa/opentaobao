package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest 服务商回传保单信息 API请求
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
type TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest struct {
	model.Params
	// 服务单投保信息回传入参
	_serviceInsuranceCallbackRequest *ServiceInsuranceCallbackRequest
}

// NewTmallSerivcecenterServicerorderInsuranceCallbackRequest 初始化TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest对象
func NewTmallSerivcecenterServicerorderInsuranceCallbackRequest() *TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest {
	return &TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest) GetApiMethodName() string {
	return "tmall.serivcecenter.servicerorder.insurance.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceInsuranceCallbackRequest is ServiceInsuranceCallbackRequest Setter
// 服务单投保信息回传入参
func (r *TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest) SetServiceInsuranceCallbackRequest(_serviceInsuranceCallbackRequest *ServiceInsuranceCallbackRequest) error {
	r._serviceInsuranceCallbackRequest = _serviceInsuranceCallbackRequest
	r.Set("service_insurance_callback_request", _serviceInsuranceCallbackRequest)
	return nil
}

// GetServiceInsuranceCallbackRequest ServiceInsuranceCallbackRequest Getter
func (r TmallSerivcecenterServicerorderInsuranceCallbackAPIRequest) GetServiceInsuranceCallbackRequest() *ServiceInsuranceCallbackRequest {
	return r._serviceInsuranceCallbackRequest
}
