package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionCreateAPIRequest 处方外流-创建处方 API请求
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
type AlibabaAlihealthOutflowPrescriptionCreateAPIRequest struct {
	model.Params
	// 保存处方入参
	_createRequest *PrescriptionOutflowUpdateRequest
}

// NewAlibabaAlihealthOutflowPrescriptionCreateRequest 初始化AlibabaAlihealthOutflowPrescriptionCreateAPIRequest对象
func NewAlibabaAlihealthOutflowPrescriptionCreateRequest() *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest {
	return &AlibabaAlihealthOutflowPrescriptionCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) Reset() {
	r._createRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.prescription.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 保存处方入参
func (r *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) SetCreateRequest(_createRequest *PrescriptionOutflowUpdateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) GetCreateRequest() *PrescriptionOutflowUpdateRequest {
	return r._createRequest
}

var poolAlibabaAlihealthOutflowPrescriptionCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowPrescriptionCreateRequest()
	},
}

// GetAlibabaAlihealthOutflowPrescriptionCreateRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowPrescriptionCreateAPIRequest
func GetAlibabaAlihealthOutflowPrescriptionCreateAPIRequest() *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest {
	return poolAlibabaAlihealthOutflowPrescriptionCreateAPIRequest.Get().(*AlibabaAlihealthOutflowPrescriptionCreateAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowPrescriptionCreateAPIRequest 将 AlibabaAlihealthOutflowPrescriptionCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowPrescriptionCreateAPIRequest(v *AlibabaAlihealthOutflowPrescriptionCreateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowPrescriptionCreateAPIRequest.Put(v)
}
