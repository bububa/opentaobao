package alihealth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthprescriptionauthgetAPIRequest 阿里健康处方平台获取授权码 API请求
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
type AlibabaalihealthprescriptionauthgetAPIRequest struct {
	model.Params
	// 请求入参
	_prescriptionRequest *PrescriptionDoctorAuthRequest
}

// NewAlibabaalihealthprescriptionauthgetRequest 初始化AlibabaalihealthprescriptionauthgetAPIRequest对象
func NewAlibabaalihealthprescriptionauthgetRequest() *AlibabaalihealthprescriptionauthgetAPIRequest {
	return &AlibabaalihealthprescriptionauthgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthprescriptionauthgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.prescription.auth.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthprescriptionauthgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthprescriptionauthgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPrescriptionRequest is PrescriptionRequest Setter
// 请求入参
func (r *AlibabaalihealthprescriptionauthgetAPIRequest) SetPrescriptionRequest(_prescriptionRequest *PrescriptionDoctorAuthRequest) error {
	r._prescriptionRequest = _prescriptionRequest
	r.Set("prescription_request", _prescriptionRequest)
	return nil
}

// GetPrescriptionRequest PrescriptionRequest Getter
func (r AlibabaalihealthprescriptionauthgetAPIRequest) GetPrescriptionRequest() *PrescriptionDoctorAuthRequest {
	return r._prescriptionRequest
}
