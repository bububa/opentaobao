package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest 阿里健康新挂号数据回传 API请求
// alibaba.alihealth.medical.registration.syncnew
//
// 阿里健康新挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalRegistrationSyncnewRequest 初始化AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncnewRequest() *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) Reset() {
	r._saveRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}

var poolAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalRegistrationSyncnewRequest()
	},
}

// GetAlibabaAlihealthMedicalRegistrationSyncnewRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest
func GetAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest() *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest {
	return poolAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest.Get().(*AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest 将 AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest(v *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalRegistrationSyncnewAPIRequest.Put(v)
}
