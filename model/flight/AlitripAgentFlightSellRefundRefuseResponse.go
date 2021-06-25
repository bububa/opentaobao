package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售退票拒绝 APIResponse
alitrip.agent.flight.sell.refund.refuse

销售退票拒绝
*/
type AlitripAgentFlightSellRefundRefuseAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellRefundRefuseResponse `json:"alitrip_agent_flight_sell_refund_refuse_response,omitempty"`
}

type AlitripAgentFlightSellRefundRefuseResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellRefundRefuseResultDto `json:"result,omitempty"`

}
