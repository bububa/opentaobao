package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签详情 APIResponse
alitrip.agent.flight.sell.modify.detail

销售改签详情
*/
type AlitripAgentFlightSellModifyDetailAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellModifyDetailResponse
}

type AlitripAgentFlightSellModifyDetailResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_detail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellModifyDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
