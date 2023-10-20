package flight

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripAgentFlightIntentionConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightIntentionConfirmAPIResponseModel).Reset()
}

// AlitripAgentFlightIntentionConfirmAPIResponseModel is 意向单确认 成功返回结果
type AlitripAgentFlightIntentionConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_intention_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightIntentionConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightIntentionConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightIntentionConfirmAPIResponse)
	},
}

// GetAlitripAgentFlightIntentionConfirmAPIResponse 从 sync.Pool 获取 AlitripAgentFlightIntentionConfirmAPIResponse
func GetAlitripAgentFlightIntentionConfirmAPIResponse() *AlitripAgentFlightIntentionConfirmAPIResponse {
	return poolAlitripAgentFlightIntentionConfirmAPIResponse.Get().(*AlitripAgentFlightIntentionConfirmAPIResponse)
}

// ReleaseAlitripAgentFlightIntentionConfirmAPIResponse 将 AlitripAgentFlightIntentionConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightIntentionConfirmAPIResponse(v *AlitripAgentFlightIntentionConfirmAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightIntentionConfirmAPIResponse.Put(v)
}
