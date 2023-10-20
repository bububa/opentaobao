package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest 处方外流-修改处方 API请求
// alibaba.alihealth.outflow.prescription.update
//
// 阿里健康-处方外流-对外提供处方修改功能
type AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest struct {
	model.Params
	// 入参对象
	_updateRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaAlihealthOutflowPrescriptionUpdateRequest 初始化AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest对象
func NewAlibabaAlihealthOutflowPrescriptionUpdateRequest() *AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest {
	return &AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) Reset() {
	r._updateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 入参对象
func (r *AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) SetUpdateRequest(_updateRequest *PrescriptionOutflowUpdateRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) GetUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return r._updateRequest
}

var poolAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowPrescriptionUpdateRequest()
	},
}

// GetAlibabaAlihealthOutflowPrescriptionUpdateRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest
func GetAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest() *AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest {
	return poolAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest.Get().(*AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest 将 AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest(v *AlibabaAlihealthOutflowPrescriptionUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionUpdateAPIRequest.Put(v)
}
