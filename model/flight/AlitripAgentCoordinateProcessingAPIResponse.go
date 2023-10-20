package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateProcessingAPIResponse 慧飞商家协同单处理完成接口 API返回值
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
type AlitripAgentCoordinateProcessingAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateProcessingAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateProcessingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateProcessingAPIResponseModel).Reset()
}

// AlitripAgentCoordinateProcessingAPIResponseModel is 慧飞商家协同单处理完成接口 成功返回结果
type AlitripAgentCoordinateProcessingAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_processing_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单处理完成接口返回结果
	Result *AlitripAgentCoordinateProcessingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateProcessingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateProcessingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateProcessingAPIResponse)
	},
}

// GetAlitripAgentCoordinateProcessingAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateProcessingAPIResponse
func GetAlitripAgentCoordinateProcessingAPIResponse() *AlitripAgentCoordinateProcessingAPIResponse {
	return poolAlitripAgentCoordinateProcessingAPIResponse.Get().(*AlitripAgentCoordinateProcessingAPIResponse)
}

// ReleaseAlitripAgentCoordinateProcessingAPIResponse 将 AlitripAgentCoordinateProcessingAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateProcessingAPIResponse(v *AlitripAgentCoordinateProcessingAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateProcessingAPIResponse.Put(v)
}
