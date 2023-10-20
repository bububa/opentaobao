package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentTicketIssueAPIResponse 【国际机票】手工出票 API返回值
// taobao.alitrip.ie.agent.ticket.issue
//
// 代理商手工出票，并回填票号
type TaobaoAlitripIeAgentTicketIssueAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentTicketIssueAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentTicketIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentTicketIssueAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentTicketIssueAPIResponseModel is 【国际机票】手工出票 成功返回结果
type TaobaoAlitripIeAgentTicketIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_ticket_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回填票号是否成功，true:成功,false：失败
	TicketSuccess bool `json:"ticket_success,omitempty" xml:"ticket_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentTicketIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TicketSuccess = false
}

var poolTaobaoAlitripIeAgentTicketIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentTicketIssueAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentTicketIssueAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentTicketIssueAPIResponse
func GetTaobaoAlitripIeAgentTicketIssueAPIResponse() *TaobaoAlitripIeAgentTicketIssueAPIResponse {
	return poolTaobaoAlitripIeAgentTicketIssueAPIResponse.Get().(*TaobaoAlitripIeAgentTicketIssueAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentTicketIssueAPIResponse 将 TaobaoAlitripIeAgentTicketIssueAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentTicketIssueAPIResponse(v *TaobaoAlitripIeAgentTicketIssueAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentTicketIssueAPIResponse.Put(v)
}
