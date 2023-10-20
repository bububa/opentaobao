package shenjing

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse 访客通过PAD提交访客码 API返回值
// alibaba.ib.shenjing.visitor.pad.fetchcodeverify
//
// 访客通过PAD提交访客码，录脸进入园区。
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse struct {
	model.CommonResponse
	AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponseModel).Reset()
}

// AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponseModel is 访客通过PAD提交访客码 成功返回结果
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_fetchcodeverify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// request_id
	ResultRequestId string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 内容
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultRequestId = ""
	m.ResultMsg = ""
	m.Content = nil
}

var poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse)
	},
}

// GetAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse
func GetAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse() *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse {
	return poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse.Get().(*AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse)
}

// ReleaseAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse 将 AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse(v *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse.Put(v)
}
