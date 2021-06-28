package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售改签详情 APIResponse
alitrip.agent.flight.sell.modify.detail

销售改签详情
*/
type AlitripAgentFlightSellModifyDetailAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellModifyDetailResponse `json:"alitrip_agent_flight_sell_modify_detail_response,omitempty"` 
    AlitripAgentFlightSellModifyDetailResponse
}

/* model for simplify = false
type AlitripAgentFlightSellModifyDetailResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        AlitripAgentFlightSellModifyDetailResultDto  *AlitripAgentFlightSellModifyDetailResultDto `json:"alitrip_agent_flight_sell_modify_detail_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellModifyDetailResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellModifyDetailResultDto `json:"result,omitempty"`

}
