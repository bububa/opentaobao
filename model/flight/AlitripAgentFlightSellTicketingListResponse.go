package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售出票列表 APIResponse
alitrip.agent.flight.sell.ticketing.list

销售出票列表
*/
type AlitripAgentFlightSellTicketingListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_agent_flight_sell_ticketing_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *PageDto `json:"result,omitempty" xml:"