package alihealthoutflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.hospital.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UpdateRequest Setter
// 入参对象
func (r *AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// Get UpdateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionHospitalVerifyAPIRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return r._updateRequest
}
