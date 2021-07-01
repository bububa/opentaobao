package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalHospitalSyncAPIRequest
阿里将康支付宝挂号数据医院回传接口 API请求
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口 */
type AlibabaAlihealthMedicalHospitalSyncAPIRequest struct {
	model.Params
	// top保存入参
	_saveRequest *CommonRequest4Top
}

// NewAlibabaAlihealthMedicalHospitalSyncRequest 初始化AlibabaAlihealthMedicalHospitalSyncAPIRequest对象
func NewAlibabaAlihealthMedicalHospitalSyncRequest() *AlibabaAlihealthMedicalHospitalSyncAPIRequest {
	return &AlibabaAlihealthMedicalHospitalSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalHospitalSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.hospital.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalHospitalSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SaveRequest Setter
// top保存入参
func (r *AlibabaAlihealthMedicalHospitalSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4Top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// Get SaveRequest Getter
func (r AlibabaAlihealthMedicalHospitalSyncAPIRequest) GetSaveRequest() *CommonRequest4Top {
	return r._saveRequest
}
