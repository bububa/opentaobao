package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse 阿里健康新挂号数据回传 API返回值
// alibaba.alihealth.medical.registration.syncnew
//
// 阿里健康新挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalRegistrationSyncnewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalRegistrationSyncnewAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalRegistrationSyncnewAPIResponseModel is 阿里健康新挂号数据回传 成功返回结果
type AlibabaAlihealthMedicalRegistrationSyncnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_registration_syncnew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalRegistrationSyncnewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse
func GetAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse() *AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse {
	return poolAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse.Get().(*AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse 将 AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse(v *AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalRegistrationSyncnewAPIResponse.Put(v)
}
