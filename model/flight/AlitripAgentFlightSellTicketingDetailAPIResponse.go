package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingDetailAPIResponse 销售出票详情 API返回值
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
type AlitripAgentFlightSellTicketingDetailAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightSellTicketingDetailAPIResponseModel
}

// AlitripAgentFlightSellTicketingDetailAPIResponseModel is 销售出票详情 成功返回结果
type AlitripAgentFlightSellTicketingDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_ticketing_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *AlitripAgentFlightSellTicketingDetailResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
