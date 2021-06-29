package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签回填 APIResponse
alitrip.agent.flight.sell.modify.backfill

销售改签回填
*/
type AlitripAgentFlightSellModifyBackfillAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellModifyBackfillResponse
}

type AlitripAgentFlightSellModifyBackfillResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_backfill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellModifyBackfillResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
