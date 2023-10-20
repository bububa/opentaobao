package alscmerchant

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlscDaodianTicketReserveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscDaodianTicketReserveAPIResponseModel).Reset()
}

// AlibabaAlscDaodianTicketReserveAPIResponseModel is 外部券冲正 成功返回结果
type AlibabaAlscDaodianTicketReserveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_daodian_ticket_reserve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlscDaodianTicketReserveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscDaodianTicketReserveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscDaodianTicketReserveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscDaodianTicketReserveAPIResponse)
	},
}

// GetAlibabaAlscDaodianTicketReserveAPIResponse 从 sync.Pool 获取 AlibabaAlscDaodianTicketReserveAPIResponse
func GetAlibabaAlscDaodianTicketReserveAPIResponse() *AlibabaAlscDaodianTicketReserveAPIResponse {
	return poolAlibabaAlscDaodianTicketReserveAPIResponse.Get().(*AlibabaAlscDaodianTicketReserveAPIResponse)
}

// ReleaseAlibabaAlscDaodianTicketReserveAPIResponse 将 AlibabaAlscDaodianTicketReserveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscDaodianTicketReserveAPIResponse(v *AlibabaAlscDaodianTicketReserveAPIResponse) {
	v.Reset()
	poolAlibabaAlscDaodianTicketReserveAPIResponse.Put(v)
}
