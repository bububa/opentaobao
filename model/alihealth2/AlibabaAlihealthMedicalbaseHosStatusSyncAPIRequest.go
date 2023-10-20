package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest 挂号医院上下线 API请求
// alibaba.alihealth.medicalbase.hos.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest struct {
	model.Params
	// 服务商医院编码
	_isvHosCode string
	// 服务项可空  默认 100001 预约挂号服务项
	_functionCode string
	// normal 正常  maintaining 维护中
	_status string
}

// NewAlibabaAlihealthMedicalbaseHosStatusSyncRequest 初始化AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseHosStatusSyncRequest() *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) Reset() {
	r._isvHosCode = ""
	r._functionCode = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.hos.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvHosCode is IsvHosCode Setter
// 服务商医院编码
func (r *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) SetIsvHosCode(_isvHosCode string) error {
	r._isvHosCode = _isvHosCode
	r.Set("isv_hos_code", _isvHosCode)
	return nil
}

// GetIsvHosCode IsvHosCode Getter
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetIsvHosCode() string {
	return r._isvHosCode
}

// SetFunctionCode is FunctionCode Setter
// 服务项可空  默认 100001 预约挂号服务项
func (r *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) SetFunctionCode(_functionCode string) error {
	r._functionCode = _functionCode
	r.Set("function_code", _functionCode)
	return nil
}

// GetFunctionCode FunctionCode Getter
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetFunctionCode() string {
	return r._functionCode
}

// SetStatus is Status Setter
// normal 正常  maintaining 维护中
func (r *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) GetStatus() string {
	return r._status
}

var poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseHosStatusSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseHosStatusSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest
func GetAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest() *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest {
	return poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest.Get().(*AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest 将 AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest(v *AlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseHosStatusSyncAPIRequest.Put(v)
}
