package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingListAPIResponse 销售出票列表 API返回值
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
type AlitripAgentFlightSellTicketingListAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellTicketingListAPIResponseModel
}

// AlitripAgentFlightSellTicketingListAPIResponseModel is 销售出票列表 成功返回结果
type AlitripAgentFlightSellTicketingListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}
