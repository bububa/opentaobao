package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundlistAPIResponse 销售退票单列表 API返回值
// alitrip.agent.flight.sell.refund.list
//
// 销售退票单列表
type AlitripagentflightsellrefundlistAPIResponse struct {
	model.CommonResponse
	AlitripagentflightsellrefundlistAPIResponseModel
}

// AlitripagentflightsellrefundlistAPIResponseModel is 销售退票单列表 成功返回结果
type AlitripagentflightsellrefundlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}
