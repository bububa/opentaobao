package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse 挂号医生上下架 API返回值
// alibaba.alihealth.medicalbase.doctor.status.sync
//
// 挂号医院上下线
type AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponseModel
}

// AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponseModel is 挂号医生上下架 成功返回结果
type AlibabaAlihealthMedicalbaseDoctorStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_doctor_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
