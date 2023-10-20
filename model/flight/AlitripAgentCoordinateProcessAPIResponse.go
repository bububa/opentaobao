package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateProcessAPIResponse 慧飞商家协同单处理完成接口 API返回值
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
type AlitripAgentCoordinateProcessAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateProcessAPIResponseModel
}

// AlitripAgentCoordinateProcessAPIResponseModel is 慧飞商家协同单处理完成接口 成功返回结果
type AlitripAgentCoordinateProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单处理完成接口返回结果
	Result *AlitripAgentCoordinateProcessResult `json:"result,omitempty" xml:"result,omitempty"`
}
