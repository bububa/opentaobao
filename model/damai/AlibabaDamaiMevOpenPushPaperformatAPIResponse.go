package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushPaperformatAPIResponse 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API返回值
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
type AlibabaDamaiMevOpenPushPaperformatAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenPushPaperformatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushPaperformatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenPushPaperformatAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenPushPaperformatAPIResponseModel is 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat 成功返回结果
type AlibabaDamaiMevOpenPushPaperformatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_push_paperformat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenPushPaperformatResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenPushPaperformatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenPushPaperformatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenPushPaperformatAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenPushPaperformatAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenPushPaperformatAPIResponse
func GetAlibabaDamaiMevOpenPushPaperformatAPIResponse() *AlibabaDamaiMevOpenPushPaperformatAPIResponse {
	return poolAlibabaDamaiMevOpenPushPaperformatAPIResponse.Get().(*AlibabaDamaiMevOpenPushPaperformatAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenPushPaperformatAPIResponse 将 AlibabaDamaiMevOpenPushPaperformatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushPaperformatAPIResponse(v *AlibabaDamaiMevOpenPushPaperformatAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushPaperformatAPIResponse.Put(v)
}
