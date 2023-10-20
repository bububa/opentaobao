package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse 乐税审核结果通知 API返回值
// alibaba.alihealth.doctor.leshui.audit.result
//
// 乐税审核结果通知
type AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDoctorLeshuiAuditResultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDoctorLeshuiAuditResultAPIResponseModel).Reset()
}

// AlibabaAlihealthDoctorLeshuiAuditResultAPIResponseModel is 乐税审核结果通知 成功返回结果
type AlibabaAlihealthDoctorLeshuiAuditResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_audit_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiAuditResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse)
	},
}

// GetAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse
func GetAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse() *AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse {
	return poolAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse.Get().(*AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse)
}

// ReleaseAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse 将 AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse(v *AlibabaAlihealthDoctorLeshuiAuditResultAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDoctorLeshuiAuditResultAPIResponse.Put(v)
}
