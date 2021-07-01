package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDoctorPublishAPIRequest
三方机构医生信息上传 API请求
alibaba.alihealth.medical.doctor.publish

三方机构上传医生信息 */
type AlibabaAlihealthMedicalDoctorPublishAPIRequest struct {
	model.Params
	// 三方机构医生上传request
	_outerDoctorPublishRequest *OuterDoctorPublishRequest
}

// New
