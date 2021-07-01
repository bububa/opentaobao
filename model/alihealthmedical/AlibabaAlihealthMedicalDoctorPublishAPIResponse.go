package alihealthmedical

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDoctorPublishAPIResponse
三方机构医生信息上传 API返回值
alibaba.alihealth.medical.doctor.publish

三方机构上传医生信息 */
type AlibabaAlihealthMedicalDoctorPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalDoctorPublishAPIResponseModel
}

// AlibabaAlihealthMedicalDoctorPublishAPIResponseModel is 三方机构医生信息上传 成功返回结果
type AlibabaAlihealthMedicalDoctorPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
