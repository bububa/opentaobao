package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售退票确认 APIResponse
alitrip.agent.flight.sell.refund.approve

销售退票确认
*/
type AlitripAgentFlightSellRefundApproveAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellRefundApproveResponse `json:"alitrip_agent_flight_sell_refund_approve_response,omitempty"` 
    AlitripAgentFlightSellRefundApproveResponse
}

/* model for simplify = false
type AlitripAgentFlightSellRefundApproveResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        AlitripAgentFlightSellRefundApproveResultDto  *AlitripAgentFlightSellRefundApproveResultDto `json:"alitrip_agent_flight_sell_refund_approve_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellRefundApproveResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellRefundApproveResultDto `json:"result,omitempty"`

}
