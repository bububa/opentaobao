package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorPublishAPIRequest 三方机构医生信息上传 API请求
// alibaba.alihealth.medical.doctor.publish
//
// 三方机构上传医生信息
type AlibabaAlihealthMedicalDoctorPublishAPIRequest struct {
	model.Params
	// 三方机构医生上传request
	_outerDoctorPublishRequest *OuterDoctorPublishRequest
}

// NewAlibabaAlihealthMedicalDoctorPublishRequest 初始化AlibabaAlihealthMedicalDoctorPublishAPIRequest对象
func NewAlibabaAlihealthMedicalDoctorPublishRequest() *AlibabaAlihealthMedicalDoctorPublishAPIRequest {
	return &AlibabaAlihealthMedicalDoctorPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalDoctorPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.doctor.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalDoctorPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterDoctorPublishRequest Setter
// 三方机构医生上传request
func (r *AlibabaAlihealthMedicalDoctorPublishAPIRequest) SetOuterDoctorPublishRequest(_outerDoctorPublishRequest *OuterDoctorPublishRequest) error {
	r._outerDoctorPublishRequest = _outerDoctorPublishRequest
	r.Set("outer_doctor_publish_request", _outerDoctorPublishRequest)
	return nil
}

// Get OuterDoctorPublishRequest Getter
func (r AlibabaAlihealthMedicalDoctorPublishAPIRequest) GetOuterDoctorPublishRequest() *OuterDoctorPublishRequest {
	return r._outerDoctorPublishRequest
}
