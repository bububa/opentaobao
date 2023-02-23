package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateListAPIResponse 慧飞商家协同单列表查询接口 API返回值
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
type AlitripAgentCoordinateListAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateListAPIResponseModel
}

// AlitripAgentCoordinateListAPIResponseModel is 慧飞商家协同单列表查询接口 成功返回结果
type AlitripAgentCoordinateListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单列表查询结果
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}
