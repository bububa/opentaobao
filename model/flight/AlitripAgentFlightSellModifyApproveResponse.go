package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签确认 APIResponse
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
type AlitripAgentFlightSellModifyApproveAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellModifyApproveResponse
}

type AlitripAgentFlightSellModifyApproveResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_approve_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellModifyApproveResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
