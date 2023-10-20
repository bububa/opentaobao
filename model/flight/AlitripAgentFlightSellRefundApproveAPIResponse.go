package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundapproveAPIResponse 销售退票确认 API返回值
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
type AlitripagentflightsellrefundapproveAPIResponse struct {
	model.CommonResponse
	AlitripagentflightsellrefundapproveAPIResponseModel
}

// AlitripagentflightsellrefundapproveAPIResponseModel is 销售退票确认 成功返回结果
type AlitripagentflightsellrefundapproveAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_approve_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripagentflightsellrefundapproveResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
