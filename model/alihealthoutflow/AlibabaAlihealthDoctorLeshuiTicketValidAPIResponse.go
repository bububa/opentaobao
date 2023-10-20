package alihealthoutflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse 乐税token验证 API返回值
// alibaba.alihealth.doctor.leshui.ticket.valid
//
// 乐税token验证
type AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDoctorLeshuiTicketValidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDoctorLeshuiTicketValidAPIResponseModel).Reset()
}

// AlibabaAlihealthDoctorLeshuiTicketValidAPIResponseModel is 乐税token验证 成功返回结果
type AlibabaAlihealthDoctorLeshuiTicketValidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_doctor_leshui_ticket_valid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDoctorLeshuiTicketValidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse)
	},
}

// GetAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse
func GetAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse() *AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse {
	return poolAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse.Get().(*AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse)
}

// ReleaseAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse 将 AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse(v *AlibabaAlihealthDoctorLeshuiTicketValidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDoctorLeshuiTicketValidAPIResponse.Put(v)
}
