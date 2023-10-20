package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncAPIResponse 阿里健康支付宝挂号记录回传接口 API返回值
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalRegistrationSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel is 阿里健康支付宝挂号记录回传接口 成功返回结果
type AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_registration_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalRegistrationSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalRegistrationSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalRegistrationSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalRegistrationSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalRegistrationSyncAPIResponse
func GetAlibabaAlihealthMedicalRegistrationSyncAPIResponse() *AlibabaAlihealthMedicalRegistrationSyncAPIResponse {
	return poolAlibabaAlihealthMedicalRegistrationSyncAPIResponse.Get().(*AlibabaAlihealthMedicalRegistrationSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalRegistrationSyncAPIResponse 将 AlibabaAlihealthMedicalRegistrationSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalRegistrationSyncAPIResponse(v *AlibabaAlihealthMedicalRegistrationSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalRegistrationSyncAPIResponse.Put(v)
}
