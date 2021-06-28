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
    AlitripAgentFlightSellTicketingListResponse
}

type AlitripAgentFlightSellTicketingListResponse struct {
    XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *PageDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
