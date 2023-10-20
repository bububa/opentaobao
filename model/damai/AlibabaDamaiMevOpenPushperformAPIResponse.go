package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushperformAPIResponse 大麦换验平台-第三方对外开放-场次接口pushPerform API返回值
// alibaba.damai.mev.open.pushperform
//
// pushPerform
type AlibabaDamaiMevOpenPushperformAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushperformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushperformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushperformAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushperformAPIResponseModel is 大麦换验平台-第三方对外开放-场次接口pushPerform 成功返回结果
type AlibabaDamaiMevOpenPushperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushperform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushperformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushperformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushperformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushperformAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushperformAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushperformAPIResponse
func GetAlibabaDamaiMevOpenPushperformAPIResponse() *AlibabaDamaiMevOpenPushperformAPIResponse {
	return poolAlibabaDamaiMevOpenPushperformAPIResponse.Get().(*AlibabaDamaiMevOpenPushperformAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushperformAPIResponse 将 AlibabaDamaiMevOpenPushperformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushperformAPIResponse(v *AlibabaDamaiMevOpenPushperformAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushperformAPIResponse.Put(v)
}
