package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售出票列表 APIResponse
alitrip.agent.flight.sell.ticketing.list

销售出票列表
*/
type AlitripAgentFlightSellTicketingListAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellTicketingListResponse `json:"alitrip_agent_flight_sell_ticketing_list_response,omitempty"` 
    AlitripAgentFlightSellTicketingListResponse
}

/* model for simplify = false
type AlitripAgentFlightSellTicketingListResponse struct {

    // 系统自动生成
    
    Result  *struct {
        PageDto  *PageDto `json:"page_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellTicketingListResponse struct {

    // 系统自动生成
    Result   *PageDto `json:"result,omitempty"`

}
