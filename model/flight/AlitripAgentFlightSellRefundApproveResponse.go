package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票确认 APIResponse
alitrip.agent.flight.sell.refund.approve

销售退票确认
*/
type AlitripAgentFlightSellRefundApproveAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellRefundApproveResponse
}

type AlitripAgentFlightSellRefundApproveResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_approve_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *AlitripAgentFlightSellRefundApproveResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
