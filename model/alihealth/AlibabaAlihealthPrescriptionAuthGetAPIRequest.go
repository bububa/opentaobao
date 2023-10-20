package alihealth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPrescriptionAuthGetAPIRequest 阿里健康处方平台获取授权码 API请求
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
type AlibabaAlihealthPrescriptionAuthGetAPIRequest struct {
	model.Params
	// 请求入参
	_prescriptionRequest *PrescriptionDoctorAuthRequest
}

// NewAlibabaAlihealthPrescriptionAuthGetRequest 初始化AlibabaAlihealthPrescriptionAuthGetAPIRequest对象
func NewAlibabaAlihealthPrescriptionAuthGetRequest() *AlibabaAlihealthPrescriptionAuthGetAPIRequest {
	return &AlibabaAlihealthPrescriptionAuthGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPrescriptionAuthGetAPIRequest) Reset() {
	r._prescriptionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPrescriptionAuthGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.prescription.auth.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPrescriptionAuthGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPrescriptionAuthGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPrescriptionRequest is PrescriptionRequest Setter
// 请求入参
func (r *AlibabaAlihealthPrescriptionAuthGetAPIRequest) SetPrescriptionRequest(_prescriptionRequest *PrescriptionDoctorAuthRequest) error {
	r._prescriptionRequest = _prescriptionRequest
	r.Set("prescription_request", _prescriptionRequest)
	return nil
}

// GetPrescriptionRequest PrescriptionRequest Getter
func (r AlibabaAlihealthPrescriptionAuthGetAPIRequest) GetPrescriptionRequest() *PrescriptionDoctorAuthRequest {
	return r._prescriptionRequest
}

var poolAlibabaAlihealthPrescriptionAuthGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPrescriptionAuthGetRequest()
	},
}

// GetAlibabaAlihealthPrescriptionAuthGetRequest 从 sync.Pool 获取 AlibabaAlihealthPrescriptionAuthGetAPIRequest
func GetAlibabaAlihealthPrescriptionAuthGetAPIRequest() *AlibabaAlihealthPrescriptionAuthGetAPIRequest {
	return poolAlibabaAlihealthPrescriptionAuthGetAPIRequest.Get().(*AlibabaAlihealthPrescriptionAuthGetAPIRequest)
}

// ReleaseAlibabaAlihealthPrescriptionAuthGetAPIRequest 将 AlibabaAlihealthPrescriptionAuthGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPrescriptionAuthGetAPIRequest(v *AlibabaAlihealthPrescriptionAuthGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPrescriptionAuthGetAPIRequest.Put(v)
}
