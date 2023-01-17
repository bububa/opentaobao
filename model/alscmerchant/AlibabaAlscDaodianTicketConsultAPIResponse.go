package alscmerchant

import (
	"encoding/xml"

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

// AlibabaAlscDaodianTicketConsultAPIResponseModel is 券码预览 成功返回结果
type AlibabaAlscDaodianTicketConsultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_daodian_ticket_consult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlscDaodianTicketConsultResult `json:"result,omitempty" xml:"result,omitempty"`
}
