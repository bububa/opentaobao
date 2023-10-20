package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushvenueAPIResponse 大麦换验平台-第三方对外开放-场馆接口pushVenue API返回值
// alibaba.damai.mev.open.pushvenue
//
// 开放接口推送场馆
type AlibabaDamaiMevOpenPushvenueAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushvenueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushvenueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushvenueAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushvenueAPIResponseModel is 大麦换验平台-第三方对外开放-场馆接口pushVenue 成功返回结果
type AlibabaDamaiMevOpenPushvenueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushvenue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushvenueResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushvenueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushvenueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushvenueAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushvenueAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushvenueAPIResponse
func GetAlibabaDamaiMevOpenPushvenueAPIResponse() *AlibabaDamaiMevOpenPushvenueAPIResponse {
	return poolAlibabaDamaiMevOpenPushvenueAPIResponse.Get().(*AlibabaDamaiMevOpenPushvenueAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushvenueAPIResponse 将 AlibabaDamaiMevOpenPushvenueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushvenueAPIResponse(v *AlibabaDamaiMevOpenPushvenueAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushvenueAPIResponse.Put(v)
}
