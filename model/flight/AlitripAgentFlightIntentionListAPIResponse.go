package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightIntentionListAPIResponse 意向单列表 API返回值
// alitrip.agent.flight.intention.list
//
// 意向单列表
type AlitripAgentFlightIntentionListAPIResponse struct {
	model.CommonResponse
	AlitripAgentFlightIntentionListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentFlightIntentionListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentFlightIntentionListAPIResponseModel).Reset()
}

// AlitripAgentFlightIntentionListAPIResponseModel is 意向单列表 成功返回结果
type AlitripAgentFlightIntentionListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_intention_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页对象
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentFlightIntentionListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentFlightIntentionListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightIntentionListAPIResponse)
	},
}

// GetAlitripAgentFlightIntentionListAPIResponse 从 sync.Pool 获取 AlitripAgentFlightIntentionListAPIResponse
func GetAlitripAgentFlightIntentionListAPIResponse() *AlitripAgentFlightIntentionListAPIResponse {
	return poolAlitripAgentFlightIntentionListAPIResponse.Get().(*AlitripAgentFlightIntentionListAPIResponse)
}

// ReleaseAlitripAgentFlightIntentionListAPIResponse 将 AlitripAgentFlightIntentionListAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentFlightIntentionListAPIResponse(v *AlitripAgentFlightIntentionListAPIResponse) {
	v.Reset()
	poolAlitripAgentFlightIntentionListAPIResponse.Put(v)
}
