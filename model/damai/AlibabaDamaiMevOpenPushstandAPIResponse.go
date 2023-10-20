package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushstandAPIResponse 大麦换验平台-第三方对外开放-看台接口pushStand API返回值
// alibaba.damai.mev.open.pushstand
//
// pushStand
type AlibabaDamaiMevOpenPushstandAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushstandAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushstandAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushstandAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushstandAPIResponseModel is 大麦换验平台-第三方对外开放-看台接口pushStand 成功返回结果
type AlibabaDamaiMevOpenPushstandAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushstand_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushstandResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushstandAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushstandAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushstandAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushstandAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushstandAPIResponse
func GetAlibabaDamaiMevOpenPushstandAPIResponse() *AlibabaDamaiMevOpenPushstandAPIResponse {
	return poolAlibabaDamaiMevOpenPushstandAPIResponse.Get().(*AlibabaDamaiMevOpenPushstandAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushstandAPIResponse 将 AlibabaDamaiMevOpenPushstandAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushstandAPIResponse(v *AlibabaDamaiMevOpenPushstandAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushstandAPIResponse.Put(v)
}
