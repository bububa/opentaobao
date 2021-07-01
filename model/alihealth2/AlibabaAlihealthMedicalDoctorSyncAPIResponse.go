package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalDoctorSyncAPIResponse
阿里健康预约挂号医生同步接口 API返回值
alibaba.alihealth.medical.doctor.sync

阿里健康预约挂号医生同步接口 */
type AlibabaAlihealthMedicalDoctorSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalDoctorSyncAPIResponseModel
}

// AlibabaAlihealthMedicalDoctorSyncAPIResponseModel is 阿里健康预约挂号医生同步接口 成功返回结果
type AlibabaAlihealthMedicalDoctorSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
