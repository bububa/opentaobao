package flight

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateDetailAPIResponse 商家协同单查询详情 API返回值
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
type AlitripAgentCoordinateDetailAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateDetailAPIResponseModel).Reset()
}

// AlitripAgentCoordinateDetailAPIResponseModel is 商家协同单查询详情 成功返回结果
type AlitripAgentCoordinateDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询协同单详情返回内容
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateDetailAPIResponse)
	},
}

// GetAlitripAgentCoordinateDetailAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateDetailAPIResponse
func GetAlitripAgentCoordinateDetailAPIResponse() *AlitripAgentCoordinateDetailAPIResponse {
	return poolAlitripAgentCoordinateDetailAPIResponse.Get().(*AlitripAgentCoordinateDetailAPIResponse)
}

// ReleaseAlitripAgentCoordinateDetailAPIResponse 将 AlitripAgentCoordinateDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateDetailAPIResponse(v *AlitripAgentCoordinateDetailAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateDetailAPIResponse.Put(v)
}
