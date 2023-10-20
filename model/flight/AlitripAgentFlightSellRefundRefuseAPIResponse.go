package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightsellrefundrefuseAPIResponse 销售退票拒绝 API返回值
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
type AlitripagentflightsellrefundrefuseAPIResponse struct {
	model.CommonResponse
	AlitripagentflightsellrefundrefuseAPIResponseModel
}

// AlitripagentflightsellrefundrefuseAPIResponseModel is 销售退票拒绝 成功返回结果
type AlitripagentflightsellrefundrefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_sell_refund_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripagentflightsellrefundrefuseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
