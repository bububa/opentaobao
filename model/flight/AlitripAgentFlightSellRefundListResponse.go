package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单列表 APIResponse
alitrip.agent.flight.sell.refund.list

销售退票单列表
*/
type AlitripAgentFlightSellRefundListAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellRefundListResponse
}

type AlitripAgentFlightSellRefundListResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *PageDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
