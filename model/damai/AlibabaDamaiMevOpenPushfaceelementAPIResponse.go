package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushfaceelementAPIResponse 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement API返回值
// alibaba.damai.mev.open.pushfaceelement
//
// pushFaceElement
type AlibabaDamaiMevOpenPushfaceelementAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushfaceelementAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfaceelementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushfaceelementAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushfaceelementAPIResponseModel is 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement 成功返回结果
type AlibabaDamaiMevOpenPushfaceelementAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushfaceelement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushfaceelementResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushfaceelementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushfaceelementAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushfaceelementAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushfaceelementAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushfaceelementAPIResponse
func GetAlibabaDamaiMevOpenPushfaceelementAPIResponse() *AlibabaDamaiMevOpenPushfaceelementAPIResponse {
	return poolAlibabaDamaiMevOpenPushfaceelementAPIResponse.Get().(*AlibabaDamaiMevOpenPushfaceelementAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushfaceelementAPIResponse 将 AlibabaDamaiMevOpenPushfaceelementAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushfaceelementAPIResponse(v *AlibabaDamaiMevOpenPushfaceelementAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushfaceelementAPIResponse.Put(v)
}
