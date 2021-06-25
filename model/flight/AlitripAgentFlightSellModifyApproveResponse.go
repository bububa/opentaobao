package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售改签确认 APIResponse
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
type AlitripAgentFlightSellModifyApproveAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellModifyApproveResponse `json:"alitrip_agent_flight_sell_modify_approve_response,omitempty"`
}

type AlitripAgentFlightSellModifyApproveResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellModifyApproveResultDto `json:"result,omitempty"`

}
