package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签拒绝 APIResponse
alitrip.agent.flight.sell.modify.refuse

销售改签拒绝
*/
type AlitripAgentFlightSellModifyRefuseAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellModifyRefuseResponse
}

type AlitripAgentFlightSellModifyRefuseResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_refuse_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellModifyRefuseResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
