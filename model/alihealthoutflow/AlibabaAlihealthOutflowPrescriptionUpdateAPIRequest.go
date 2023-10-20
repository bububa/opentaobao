package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowprescriptionupdateAPIRequest 处方外流-修改处方 API请求
// alibaba.alihealth.outflow.prescription.update
//
// 阿里健康-处方外流-对外提供处方修改功能
type AlibabaalihealthoutflowprescriptionupdateAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaalihealthoutflowprescriptionupdateRequest 初始化AlibabaalihealthoutflowprescriptionupdateAPIRequest对象
func NewAlibabaalihealthoutflowprescriptionupdateRequest() *AlibabaalihealthoutflowprescriptionupdateAPIRequest {
	return &AlibabaalihealthoutflowprescriptionupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowprescriptionupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowprescriptionupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowprescriptionupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 入参对象
func (r *AlibabaalihealthoutflowprescriptionupdateAPIRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaalihealthoutflowprescriptionupdateAPIRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return r._updateRequest
}
