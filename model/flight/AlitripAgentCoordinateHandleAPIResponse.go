package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinatehandleAPIResponse 慧飞商家协同单接手接口 API返回值
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
type AlitripagentcoordinatehandleAPIResponse struct {
	model.CommonResponse
	AlitripagentcoordinatehandleAPIResponseModel
}

// AlitripagentcoordinatehandleAPIResponseModel is 慧飞商家协同单接手接口 成功返回结果
type AlitripagentcoordinatehandleAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_handle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单接手接口返回结果
	Result *AlitripagentcoordinatehandleResult `json:"result,omitempty" xml:"result,omitempty"`
}
