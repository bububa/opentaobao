package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售改签单列表 APIResponse
alitrip.agent.flight.sell.modify.list

销售改签单列表
*/
type AlitripAgentFlightSellModifyListAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellModifyListResponse `json:"alitrip_agent_flight_sell_modify_list_response,omitempty"`
}

type AlitripAgentFlightSellModifyListResponse struct {

    // 系统自动生成
    Result   *PageDto `json:"result,omitempty"`

}
