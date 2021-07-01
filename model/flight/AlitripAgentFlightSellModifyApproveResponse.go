package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售改签确认 API返回值 
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
type AlitripAgentFlightSellModifyApproveAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellModifyApproveResponse
}

// 销售改签确认 成功返回结果
type AlitripAgentFlightSellModifyApproveResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_modify_approve_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
