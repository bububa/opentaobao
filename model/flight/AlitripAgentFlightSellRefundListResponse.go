package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单列表 APIResponse
alitrip.agent.flight.sell.refund.list

销售退票单列表
*/
type AlitripAgentFlightSellRefundListAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellRefundListResponse `json:"alitrip_agent_flight_sell_refund_list_response,omitempty"` 
    AlitripAgentFlightSellRefundListResponse
}

/* model for simplify = false
type AlitripAgentFlightSellRefundListResponse struct {

    // 系统自动生成
    
    Result  *struct {
        PageDto  *PageDto `json:"page_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellRefundListResponse struct {

    // 系统自动生成
    Result   *PageDto `json:"result,omitempty"`

}
