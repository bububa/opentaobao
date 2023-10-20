package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest 处方同步至医院返回校验结果 API请求
// alibaba.alihealth.outflow.prescription.hospital.verify
//
// 处方同步至医院返回校验结果
type AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaalihealthoutflowprescriptionhospitalverifyRequest 初始化AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest对象
func NewAlibabaalihealthoutflowprescriptionhospitalverifyRequest() *AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest {
	return &AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.hospital.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 入参对象
func (r *AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaalihealthoutflowprescriptionhospitalverifyAPIRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return r._updateRequest
}
