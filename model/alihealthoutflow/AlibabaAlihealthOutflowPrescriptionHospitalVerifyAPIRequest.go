package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest 处方同步至医院返回校验结果 API请求
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
type AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest 初始化AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest对象
func NewAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest() *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest {
	return &AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) Reset() {
	r._updateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.hospital.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 入参对象
func (r *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return r._updateRequest
}

var poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest()
	},
}

// GetAlibabaAlihealthOutflowPrescriptionHospitalVerifyRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest
func GetAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest() *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest {
	return poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest.Get().(*AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest 将 AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest(v *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest.Put(v)
}
