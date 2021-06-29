package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单详情 API返回值 
alitrip.agent.flight.sell.refund.detail

销售退票单详情
*/
type AlitripAgentFlightSellRefundDetailAPIResponse struct {
    model.CommonResponse
    AlitripAgentFlightSellRefundDetailResponse
}

// 销售退票单详情 成功返回结果
type AlitripAgentFlightSellRefundDetailResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_detail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellRefundDetailResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
