package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushitemAPIResponse 大麦换验平台-第三方对外开放-票品接口pushItem API返回值
// alibaba.damai.mev.open.pushitem
//
// 开放接口 推送票品
type AlibabaDamaiMevOpenPushitemAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushitemAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushitemAPIResponseModel is 大麦换验平台-第三方对外开放-票品接口pushItem 成功返回结果
type AlibabaDamaiMevOpenPushitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushitemResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushitemAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushitemAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushitemAPIResponse
func GetAlibabaDamaiMevOpenPushitemAPIResponse() *AlibabaDamaiMevOpenPushitemAPIResponse {
	return poolAlibabaDamaiMevOpenPushitemAPIResponse.Get().(*AlibabaDamaiMevOpenPushitemAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushitemAPIResponse 将 AlibabaDamaiMevOpenPushitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushitemAPIResponse(v *AlibabaDamaiMevOpenPushitemAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushitemAPIResponse.Put(v)
}
