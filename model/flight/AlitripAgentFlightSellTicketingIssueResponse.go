package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票 APIResponse
alitrip.agent.flight.sell.ticketing.issue

销售出票
*/
type AlitripAgentFlightSellTicketingIssueAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellTicketingIssueResponse
}

type AlitripAgentFlightSellTicketingIssueResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_issue_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellTicketingIssueResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
