package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingIssueAPIResponse 销售出票 API返回值
// alitrip.agent.flight.sell.ticketing.issue
//
// 销售出票
type AlitripAgentFlightSellTicketingIssueAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellTicketingIssueAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellTicketingIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightSellTicketingIssueAPIResponseModel).Reset()
}

// AlitripAgentFlightSellTicketingIssueAPIResponseModel is 销售出票 成功返回结果
type AlitripAgentFlightSellTicketingIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripAgentFlightSellTicketingIssueResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightSellTicketingIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightSellTicketingIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellTicketingIssueAPIResponse)
	},
}

// GetAlitripAgentFlightSellTicketingIssueAPIResponse 从 sync.Pool 获取 AlitripAgentFlightSellTicketingIssueAPIResponse
func GetAlitripAgentFlightSellTicketingIssueAPIResponse() *AlitripAgentFlightSellTicketingIssueAPIResponse {
	return poolAlitripAgentFlightSellTicketingIssueAPIResponse.Get().(*AlitripAgentFlightSellTicketingIssueAPIResponse)
}

// ReleaseAlitripAgentFlightSellTicketingIssueAPIResponse 将 AlitripAgentFlightSellTicketingIssueAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightSellTicketingIssueAPIResponse(v *AlitripAgentFlightSellTicketingIssueAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightSellTicketingIssueAPIResponse.Put(v)
}
