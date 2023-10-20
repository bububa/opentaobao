package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightIntentionConfirmAPIResponse 意向单确认 API返回值
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
type AlitripAgentFlightIntentionConfirmAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightIntentionConfirmAPIResponseModel
}

// AlitripAgentFlightIntentionConfirmAPIResponseModel is 意向单确认 成功返回结果
type AlitripAgentFlightIntentionConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_intention_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
