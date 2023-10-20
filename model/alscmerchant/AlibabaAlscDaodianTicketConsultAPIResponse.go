package alscmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscDaodianTicketConsultAPIResponse 券码预览 API返回值
// alibaba.alsc.daodian.ticket.consult
//
// 券码预览
type AlibabaAlscDaodianTicketConsultAPIResponse struct {
	model.CommonResponse
	AlibabaAlscDaodianTicketConsultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscDaodianTicketConsultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscDaodianTicketConsultAPIResponseModel).Reset()
}

// AlibabaAlscDaodianTicketConsultAPIResponseModel is 券码预览 成功返回结果
type AlibabaAlscDaodianTicketConsultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_daodian_ticket_consult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlscDaodianTicketConsultResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscDaodianTicketConsultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscDaodianTicketConsultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscDaodianTicketConsultAPIResponse)
	},
}

// GetAlibabaAlscDaodianTicketConsultAPIResponse 从 sync.Pool 获取 AlibabaAlscDaodianTicketConsultAPIResponse
func GetAlibabaAlscDaodianTicketConsultAPIResponse() *AlibabaAlscDaodianTicketConsultAPIResponse {
	return poolAlibabaAlscDaodianTicketConsultAPIResponse.Get().(*AlibabaAlscDaodianTicketConsultAPIResponse)
}

// ReleaseAlibabaAlscDaodianTicketConsultAPIResponse 将 AlibabaAlscDaodianTicketConsultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscDaodianTicketConsultAPIResponse(v *AlibabaAlscDaodianTicketConsultAPIResponse) {
	v.Reset()
	poolAlibabaAlscDaodianTicketConsultAPIResponse.Put(v)
}
