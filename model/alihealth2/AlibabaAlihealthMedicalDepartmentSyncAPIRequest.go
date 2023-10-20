package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDepartmentSyncAPIRequest 阿里健康预约挂号科室同步接口 API请求
// alibaba.alihealth.medical.department.sync
//
// 阿里健康预约挂号科室同步接口
type AlibabaAlihealthMedicalDepartmentSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalDepartmentSyncRequest 初始化AlibabaAlihealthMedicalDepartmentSyncAPIRequest对象
func NewAlibabaAlihealthMedicalDepartmentSyncRequest() *AlibabaAlihealthMedicalDepartmentSyncAPIRequest {
	return &AlibabaAlihealthMedicalDepartmentSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalDepartmentSyncAPIRequest) Reset() {
	r._saveRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDepartmentSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.department.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDepartmentSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalDepartmentSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalDepartmentSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalDepartmentSyncAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}

var poolAlibabaAlihealthMedicalDepartmentSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalDepartmentSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalDepartmentSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalDepartmentSyncAPIRequest
func GetAlibabaAlihealthMedicalDepartmentSyncAPIRequest() *AlibabaAlihealthMedicalDepartmentSyncAPIRequest {
	return poolAlibabaAlihealthMedicalDepartmentSyncAPIRequest.Get().(*AlibabaAlihealthMedicalDepartmentSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalDepartmentSyncAPIRequest 将 AlibabaAlihealthMedicalDepartmentSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalDepartmentSyncAPIRequest(v *AlibabaAlihealthMedicalDepartmentSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalDepartmentSyncAPIRequest.Put(v)
}
