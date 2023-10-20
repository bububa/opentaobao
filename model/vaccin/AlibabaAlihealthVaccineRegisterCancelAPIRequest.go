package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccineregistercancelAPIRequest 取消登记 API请求
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
type AlibabaalihealthvaccineregistercancelAPIRequest struct {
	model.Params
	// 无
	_topRequest *CancelVcRegisterRequest
}

// NewAlibabaalihealthvaccineregistercancelRequest 初始化AlibabaalihealthvaccineregistercancelAPIRequest对象
func NewAlibabaalihealthvaccineregistercancelRequest() *AlibabaalihealthvaccineregistercancelAPIRequest {
	return &AlibabaalihealthvaccineregistercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthvaccineregistercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.register.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthvaccineregistercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthvaccineregistercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRequest is TopRequest Setter
// 无
func (r *AlibabaalihealthvaccineregistercancelAPIRequest) SetTopRequest(_topRequest *CancelVcRegisterRequest) error {
	r._topRequest = _topRequest
	r.Set("top_request", _topRequest)
	return nil
}

// GetTopRequest TopRequest Getter
func (r AlibabaalihealthvaccineregistercancelAPIRequest) GetTopRequest() *CancelVcRegisterRequest {
	return r._topRequest
}
