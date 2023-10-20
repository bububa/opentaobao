package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateGobackAPIResponse 协同单驳回 API返回值
// alitrip.agent.coordinate.goback
//
// 协同单驳回
type AlitripAgentCoordinateGobackAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateGobackAPIResponseModel
}

// AlitripAgentCoordinateGobackAPIResponseModel is 协同单驳回 成功返回结果
type AlitripAgentCoordinateGobackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_goback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单驳回返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
