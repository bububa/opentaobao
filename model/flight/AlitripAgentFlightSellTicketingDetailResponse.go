package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售出票详情 APIResponse
alitrip.agent.flight.sell.ticketing.detail

销售出票详情
*/
type AlitripAgentFlightSellTicketingDetailAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellTicketingDetailResponse `json:"alitrip_agent_flight_sell_ticketing_detail_response,omitempty"`
}

type AlitripAgentFlightSellTicketingDetailResponse struct {

    // 出参对象
    Result   *AlitripAgentFlightSellTicketingDetailResultDto `json:"result,omitempty"`

}
