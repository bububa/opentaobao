package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalDoctorSyncAPIResponse 阿里健康预约挂号医生同步接口 API返回值
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
type AlibabaAlihealthMedicalDoctorSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalDoctorSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDoctorSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalDoctorSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalDoctorSyncAPIResponseModel is 阿里健康预约挂号医生同步接口 成功返回结果
type AlibabaAlihealthMedicalDoctorSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_doctor_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalDoctorSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalDoctorSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalDoctorSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalDoctorSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalDoctorSyncAPIResponse
func GetAlibabaAlihealthMedicalDoctorSyncAPIResponse() *AlibabaAlihealthMedicalDoctorSyncAPIResponse {
	return poolAlibabaAlihealthMedicalDoctorSyncAPIResponse.Get().(*AlibabaAlihealthMedicalDoctorSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalDoctorSyncAPIResponse 将 AlibabaAlihealthMedicalDoctorSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalDoctorSyncAPIResponse(v *AlibabaAlihealthMedicalDoctorSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalDoctorSyncAPIResponse.Put(v)
}
