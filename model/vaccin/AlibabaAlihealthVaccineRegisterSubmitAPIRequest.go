package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccineregistersubmitAPIRequest cdc回传疫苗登记数据 API请求
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
type AlibabaalihealthvaccineregistersubmitAPIRequest struct {
	model.Params
	// 无
	_topRequest *SubmitVcRegisterRequest
}

// NewAlibabaalihealthvaccineregistersubmitRequest 初始化AlibabaalihealthvaccineregistersubmitAPIRequest对象
func NewAlibabaalihealthvaccineregistersubmitRequest() *AlibabaalihealthvaccineregistersubmitAPIRequest {
	return &AlibabaalihealthvaccineregistersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthvaccineregistersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.register.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthvaccineregistersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthvaccineregistersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopRequest is TopRequest Setter
// 无
func (r *AlibabaalihealthvaccineregistersubmitAPIRequest) SetTopRequest(_topRequest *SubmitVcRegisterRequest) error {
	r._topRequest = _topRequest
	r.Set("top_request", _topRequest)
	return nil
}

// GetTopRequest TopRequest Getter
func (r AlibabaalihealthvaccineregistersubmitAPIRequest) GetTopRequest() *SubmitVcRegisterRequest {
	return r._topRequest
}
