package flight

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripAgentCoordinateListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripAgentCoordinateListAPIResponseModel).Reset()
}

// AlitripAgentCoordinateListAPIResponseModel is 慧飞商家协同单列表查询接口 成功返回结果
type AlitripAgentCoordinateListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单列表查询结果
	Result *PageDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripAgentCoordinateListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripAgentCoordinateListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripAgentCoordinateListAPIResponse)
	},
}

// GetAlitripAgentCoordinateListAPIResponse 从 sync.Pool 获取 AlitripAgentCoordinateListAPIResponse
func GetAlitripAgentCoordinateListAPIResponse() *AlitripAgentCoordinateListAPIResponse {
	return poolAlitripAgentCoordinateListAPIResponse.Get().(*AlitripAgentCoordinateListAPIResponse)
}

// ReleaseAlitripAgentCoordinateListAPIResponse 将 AlitripAgentCoordinateListAPIResponse 保存到 sync.Pool
func ReleaseAlitripAgentCoordinateListAPIResponse(v *AlitripAgentCoordinateListAPIResponse) {
	v.Reset()
	poolAlitripAgentCoordinateListAPIResponse.Put(v)
}
