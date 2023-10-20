package flight

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripAgentCoordinateProcessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateProcessAPIResponseModel).Reset()
}

// AlitripAgentCoordinateProcessAPIResponseModel is 慧飞商家协同单处理完成接口 成功返回结果
type AlitripAgentCoordinateProcessAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单处理完成接口返回结果
	Result *AlitripAgentCoordinateProcessResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateProcessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateProcessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateProcessAPIResponse)
	},
}

// GetAlitripAgentCoordinateProcessAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateProcessAPIResponse
func GetAlitripAgentCoordinateProcessAPIResponse() *AlitripAgentCoordinateProcessAPIResponse {
	return poolAlitripAgentCoordinateProcessAPIResponse.Get().(*AlitripAgentCoordinateProcessAPIResponse)
}

// ReleaseAlitripAgentCoordinateProcessAPIResponse 将 AlitripAgentCoordinateProcessAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateProcessAPIResponse(v *AlitripAgentCoordinateProcessAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateProcessAPIResponse.Put(v)
}
