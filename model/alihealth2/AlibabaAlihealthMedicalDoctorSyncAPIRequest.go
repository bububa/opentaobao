package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorSyncAPIRequest 阿里健康预约挂号医生同步接口 API请求
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
type AlibabaAlihealthMedicalDoctorSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaAlihealthMedicalDoctorSyncRequest 初始化AlibabaAlihealthMedicalDoctorSyncAPIRequest对象
func NewAlibabaAlihealthMedicalDoctorSyncRequest() *AlibabaAlihealthMedicalDoctorSyncAPIRequest {
	return &AlibabaAlihealthMedicalDoctorSyncAPIRequest{
		Params: model.NewParams(),
	}
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
func (r *AlibabaAlihealthMedicalDoctorSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalDoctorSyncAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
