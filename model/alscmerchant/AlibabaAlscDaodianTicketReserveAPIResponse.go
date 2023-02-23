package alscmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscDaodianTicketReserveAPIResponse 外部券冲正 API返回值
// alibaba.alsc.daodian.ticket.reserve
//
// 外部券冲正
type AlibabaAlscDaodianTicketReserveAPIResponse struct {
	model.CommonResponse
	AlibabaAlscDaodianTicketReserveAPIResponseModel
}

// AlibabaAlscDaodianTicketReserveAPIResponseModel is 外部券冲正 成功返回结果
type AlibabaAlscDaodianTicketReserveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_daodian_ticket_reserve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlscDaodianTicketReserveResult `json:"result,omitempty" xml:"result,omitempty"`
}
