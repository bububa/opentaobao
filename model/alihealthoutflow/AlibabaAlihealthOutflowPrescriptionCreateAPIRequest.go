package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthoutflowprescriptioncreateAPIRequest 处方外流-创建处方 API请求
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
type AlibabaalihealthoutflowprescriptioncreateAPIRequest struct {
	model.Params
	// 保存处方入参
	_createRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaalihealthoutflowprescriptioncreateRequest 初始化AlibabaalihealthoutflowprescriptioncreateAPIRequest对象
func NewAlibabaalihealthoutflowprescriptioncreateRequest() *AlibabaalihealthoutflowprescriptioncreateAPIRequest {
	return &AlibabaalihealthoutflowprescriptioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthoutflowprescriptioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthoutflowprescriptioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthoutflowprescriptioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 保存处方入参
func (r *AlibabaalihealthoutflowprescriptioncreateAPIRequest) SetCreateRequest(_createRequest *PrescriptionOutflowUpdateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabaalihealthoutflowprescriptioncreateAPIRequest) GetCreateRequest() *PrescriptionOutflowUpdateRequest {
	return r._createRequest
}
