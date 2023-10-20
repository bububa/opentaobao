package alihealth2

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel is 直连医生上传 成功返回结果
type AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_doctor_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse
func GetAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse() *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse {
	return poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse.Get().(*AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse 将 AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse(v *AlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseDoctorSyncnewAPIResponse.Put(v)
}
