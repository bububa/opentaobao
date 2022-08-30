package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse 直连医生上传 API返回值
// alibaba.alihealth.medicalbase.doctor.syncnew
//
// 直连医生上传
type AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel
}

// AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel is 直连医生上传 成功返回结果
type AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_doctor_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
