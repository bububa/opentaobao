package shenjing

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse pad获取二维码 API返回值
// alibaba.ib.shenjing.visitor.pad.getqrcodelink
//
// pad获取二维码链接。扫码录入人脸。
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse struct {
	model.CommonResponse
	AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponseModel).Reset()
}

// AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponseModel is pad获取二维码 成功返回结果
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ib_shenjing_visitor_pad_getqrcodelink_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 等级
	ResultLevel string `json:"result_level,omitempty" xml:"result_level,omitempty"`
	// request_id
	ResultRequestId string `json:"result_request_id,omitempty" xml:"result_request_id,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.ResultLevel = ""
	m.ResultRequestId = ""
	m.ResultSuccess = false
}

var poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse)
	},
}

// GetAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse
func GetAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse() *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse {
	return poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse.Get().(*AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse)
}

// ReleaseAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse 将 AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse(v *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse.Put(v)
}
