package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfloorAPIResponse 大麦换验平台-第三方对外开放-楼层接口pushFloor API返回值
// alibaba.damai.mev.open.pushfloor
//
// pushFloor
type AlibabaDamaiMevOpenPushfloorAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushfloorAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfloorAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushfloorAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushfloorAPIResponseModel is 大麦换验平台-第三方对外开放-楼层接口pushFloor 成功返回结果
type AlibabaDamaiMevOpenPushfloorAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushfloor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushfloorResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfloorAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushfloorAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfloorAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushfloorAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfloorAPIResponse
func GetAlibabaDamaiMevOpenPushfloorAPIResponse() *AlibabaDamaiMevOpenPushfloorAPIResponse {
	return poolAlibabaDamaiMevOpenPushfloorAPIResponse.Get().(*AlibabaDamaiMevOpenPushfloorAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushfloorAPIResponse 将 AlibabaDamaiMevOpenPushfloorAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfloorAPIResponse(v *AlibabaDamaiMevOpenPushfloorAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfloorAPIResponse.Put(v)
}
