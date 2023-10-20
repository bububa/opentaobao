package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse 获取药品扫码落地页vivo API返回值
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
type AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponseModel is 获取药品扫码落地页vivo 成功返回结果
type AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodesearch_getshowurl_vivo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 落地页地址
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse
func GetAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse() *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse {
	return poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse.Get().(*AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse 将 AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse(v *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse.Put(v)
}
