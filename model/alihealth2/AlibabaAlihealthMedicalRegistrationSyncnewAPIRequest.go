package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest
阿里健康新挂号数据回传 API请求
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口 */
type AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalRegistrationSyncnewRequest 初始化AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncnewRequest() *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// Get SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}
