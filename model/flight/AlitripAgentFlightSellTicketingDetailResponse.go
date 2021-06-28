package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票详情 APIResponse
alitrip.agent.flight.sell.ticketing.detail

销售出票详情
*/
type AlitripAgentFlightSellTicketingDetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_agent_flight_sell_ticketing_detail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参对象
    
    Result   *AlitripAgentFlightSellTicketingDetailResultDto `json:"result,omitempty" xml:"