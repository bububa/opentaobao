package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateHandleAPIResponse 慧飞商家协同单接手接口 API返回值
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
type AlitripAgentCoordinateHandleAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateHandleAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateHandleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateHandleAPIResponseModel).Reset()
}

// AlitripAgentCoordinateHandleAPIResponseModel is 慧飞商家协同单接手接口 成功返回结果
type AlitripAgentCoordinateHandleAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_handle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单接手接口返回结果
	Result *AlitripAgentCoordinateHandleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateHandleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateHandleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateHandleAPIResponse)
	},
}

// GetAlitripAgentCoordinateHandleAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateHandleAPIResponse
func GetAlitripAgentCoordinateHandleAPIResponse() *AlitripAgentCoordinateHandleAPIResponse {
	return poolAlitripAgentCoordinateHandleAPIResponse.Get().(*AlitripAgentCoordinateHandleAPIResponse)
}

// ReleaseAlitripAgentCoordinateHandleAPIResponse 将 AlitripAgentCoordinateHandleAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateHandleAPIResponse(v *AlitripAgentCoordinateHandleAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateHandleAPIResponse.Put(v)
}
