package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncAPIRequest 阿里健康支付宝挂号记录回传接口 API请求
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalRegistrationSyncRequest 初始化AlibabaAlihealthMedicalRegistrationSyncAPIRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncRequest() *AlibabaAlihealthMedicalRegistrationSyncAPIRequest {
	return &AlibabaAlihealthMedicalRegistrationSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalRegistrationSyncAPIRequest) Reset() {
	r._saveRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}

var poolAlibabaAlihealthMedicalRegistrationSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalRegistrationSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalRegistrationSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalRegistrationSyncAPIRequest
func GetAlibabaAlihealthMedicalRegistrationSyncAPIRequest() *AlibabaAlihealthMedicalRegistrationSyncAPIRequest {
	return poolAlibabaAlihealthMedicalRegistrationSyncAPIRequest.Get().(*AlibabaAlihealthMedicalRegistrationSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalRegistrationSyncAPIRequest 将 AlibabaAlihealthMedicalRegistrationSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalRegistrationSyncAPIRequest(v *AlibabaAlihealthMedicalRegistrationSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalRegistrationSyncAPIRequest.Put(v)
}
