package flight

import (
	"encoding/xml"

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

// AlitripAgentFlightIntentionListAPIResponseModel is 意向单列表 成功返回结果
type AlitripAgentFlightIntentionListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_flight_intention_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页对象
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}
