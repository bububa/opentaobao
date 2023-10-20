package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfaceAPIResponse 大麦换验平台-第三方对外开放-票面接口pushFace API返回值
// alibaba.damai.mev.open.pushface
//
// pushFace
type AlibabaDamaiMevOpenPushfaceAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushfaceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushfaceAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushfaceAPIResponseModel is 大麦换验平台-第三方对外开放-票面接口pushFace 成功返回结果
type AlibabaDamaiMevOpenPushfaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushfaceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushfaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfaceAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushfaceAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfaceAPIResponse
func GetAlibabaDamaiMevOpenPushfaceAPIResponse() *AlibabaDamaiMevOpenPushfaceAPIResponse {
	return poolAlibabaDamaiMevOpenPushfaceAPIResponse.Get().(*AlibabaDamaiMevOpenPushfaceAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushfaceAPIResponse 将 AlibabaDamaiMevOpenPushfaceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfaceAPIResponse(v *AlibabaDamaiMevOpenPushfaceAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfaceAPIResponse.Put(v)
}
