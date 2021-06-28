package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售出票 APIResponse
alitrip.agent.flight.sell.ticketing.issue

销售出票
*/
type AlitripAgentFlightSellTicketingIssueAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellTicketingIssueResponse `json:"alitrip_agent_flight_sell_ticketing_issue_response,omitempty"` 
    AlitripAgentFlightSellTicketingIssueResponse
}

/* model for simplify = false
type AlitripAgentFlightSellTicketingIssueResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        AlitripAgentFlightSellTicketingIssueResultDto  *AlitripAgentFlightSellTicketingIssueResultDto `json:"alitrip_agent_flight_sell_ticketing_issue_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellTicketingIssueResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellTicketingIssueResultDto `json:"result,omitempty"`

}
