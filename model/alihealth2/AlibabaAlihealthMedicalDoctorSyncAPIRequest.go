package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorSyncAPIRequest 阿里健康预约挂号医生同步接口 API请求
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
type AlibabaAlihealthMedicalDoctorSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalDoctorSyncRequest 初始化AlibabaAlihealthMedicalDoctorSyncAPIRequest对象
func NewAlibabaAlihealthMedicalDoctorSyncRequest() *AlibabaAlihealthMedicalDoctorSyncAPIRequest {
	return &AlibabaAlihealthMedicalDoctorSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalDoctorSyncAPIRequest) Reset() {
	r._saveRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalDoctorSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}

var poolAlibabaAlihealthMedicalDoctorSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalDoctorSyncRequest()
	},
}

// GetAlibabaAlihealthMedicalDoctorSyncRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalDoctorSyncAPIRequest
func GetAlibabaAlihealthMedicalDoctorSyncAPIRequest() *AlibabaAlihealthMedicalDoctorSyncAPIRequest {
	return poolAlibabaAlihealthMedicalDoctorSyncAPIRequest.Get().(*AlibabaAlihealthMedicalDoctorSyncAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalDoctorSyncAPIRequest 将 AlibabaAlihealthMedicalDoctorSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalDoctorSyncAPIRequest(v *AlibabaAlihealthMedicalDoctorSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalDoctorSyncAPIRequest.Put(v)
}
