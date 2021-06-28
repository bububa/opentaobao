package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售改签回填 APIResponse
alitrip.agent.flight.sell.modify.backfill

销售改签回填
*/
type AlitripAgentFlightSellModifyBackfillAPIResponse struct {
    model.CommonResponse
    // Response *AlitripAgentFlightSellModifyBackfillResponse `json:"alitrip_agent_flight_sell_modify_backfill_response,omitempty"` 
    AlitripAgentFlightSellModifyBackfillResponse
}

/* model for simplify = false
type AlitripAgentFlightSellModifyBackfillResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        AlitripAgentFlightSellModifyBackfillResultDto  *AlitripAgentFlightSellModifyBackfillResultDto `json:"alitrip_agent_flight_sell_modify_backfill_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripAgentFlightSellModifyBackfillResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellModifyBackfillResultDto `json:"result,omitempty"`

}
