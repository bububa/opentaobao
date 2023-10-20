package alitripbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBpChannelCrowQueryAPIResponse 人群匹配 API返回值
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
type AlitripBpChannelCrowQueryAPIResponse struct {
	model.CommonResponse
	AlitripBpChannelCrowQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBpChannelCrowQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBpChannelCrowQueryAPIResponseModel).Reset()
}

// AlitripBpChannelCrowQueryAPIResponseModel is 人群匹配 成功返回结果
type AlitripBpChannelCrowQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bp_channel_crow_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AdResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBpChannelCrowQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBpChannelCrowQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBpChannelCrowQueryAPIResponse)
	},
}

// GetAlitripBpChannelCrowQueryAPIResponse 从 sync.Pool 获取 AlitripBpChannelCrowQueryAPIResponse
func GetAlitripBpChannelCrowQueryAPIResponse() *AlitripBpChannelCrowQueryAPIResponse {
	return poolAlitripBpChannelCrowQueryAPIResponse.Get().(*AlitripBpChannelCrowQueryAPIResponse)
}

// ReleaseAlitripBpChannelCrowQueryAPIResponse 将 AlitripBpChannelCrowQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBpChannelCrowQueryAPIResponse(v *AlitripBpChannelCrowQueryAPIResponse) {
	v.Reset()
	poolAlitripBpChannelCrowQueryAPIResponse.Put(v)
}
