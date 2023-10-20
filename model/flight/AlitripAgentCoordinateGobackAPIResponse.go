package flight

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripAgentCoordinateGobackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateGobackAPIResponseModel).Reset()
}

// AlitripAgentCoordinateGobackAPIResponseModel is 协同单驳回 成功返回结果
type AlitripAgentCoordinateGobackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_goback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单驳回返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateGobackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateGobackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateGobackAPIResponse)
	},
}

// GetAlitripAgentCoordinateGobackAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateGobackAPIResponse
func GetAlitripAgentCoordinateGobackAPIResponse() *AlitripAgentCoordinateGobackAPIResponse {
	return poolAlitripAgentCoordinateGobackAPIResponse.Get().(*AlitripAgentCoordinateGobackAPIResponse)
}

// ReleaseAlitripAgentCoordinateGobackAPIResponse 将 AlitripAgentCoordinateGobackAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateGobackAPIResponse(v *AlitripAgentCoordinateGobackAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateGobackAPIResponse.Put(v)
}
