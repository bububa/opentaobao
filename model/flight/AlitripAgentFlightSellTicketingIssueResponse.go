package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票 API返回值 
alitrip.agent.flight.sell.ticketing.issue

销售出票
*/
type AlitripAgentFlightSellTicketingIssueAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellTicketingIssueResponse
}

// 销售出票 成功返回结果
type AlitripAgentFlightSellTicketingIssueResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_issue_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellTicketingIssueResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
