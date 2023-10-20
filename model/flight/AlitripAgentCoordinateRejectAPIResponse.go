package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateRejectAPIResponse 慧飞商家协同单拒绝接口 API返回值
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
type AlitripAgentCoordinateRejectAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateRejectAPIResponseModel
}

// AlitripAgentCoordinateRejectAPIResponseModel is 慧飞商家协同单拒绝接口 成功返回结果
type AlitripAgentCoordinateRejectAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_reject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单拒绝接口返回结果
	Result *AlitripAgentCoordinateRejectResult `json:"result,omitempty" xml:"result,omitempty"`
}
