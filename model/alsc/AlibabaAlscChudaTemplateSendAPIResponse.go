package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscChudaTemplateSendAPIResponse 本地生活触达模板消息发送接口 API返回值
// alibaba.alsc.chuda.template.send
//
// 允许三方小程序通过该api发送本地生活触达消息
type AlibabaAlscChudaTemplateSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlscChudaTemplateSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscChudaTemplateSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscChudaTemplateSendAPIResponseModel).Reset()
}

// AlibabaAlscChudaTemplateSendAPIResponseModel is 本地生活触达模板消息发送接口 成功返回结果
type AlibabaAlscChudaTemplateSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_chuda_template_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码，200表示成功
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 描述信息
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// msgId
	ResultObj int64 `json:"result_obj,omitempty" xml:"result_obj,omitempty"`
	// 发送是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscChudaTemplateSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultDesc = ""
	m.ResultObj = 0
	m.IsSuccess = false
}

var poolAlibabaAlscChudaTemplateSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscChudaTemplateSendAPIResponse)
	},
}

// GetAlibabaAlscChudaTemplateSendAPIResponse 从 sync.Pool 获取 AlibabaAlscChudaTemplateSendAPIResponse
func GetAlibabaAlscChudaTemplateSendAPIResponse() *AlibabaAlscChudaTemplateSendAPIResponse {
	return poolAlibabaAlscChudaTemplateSendAPIResponse.Get().(*AlibabaAlscChudaTemplateSendAPIResponse)
}

// ReleaseAlibabaAlscChudaTemplateSendAPIResponse 将 AlibabaAlscChudaTemplateSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscChudaTemplateSendAPIResponse(v *AlibabaAlscChudaTemplateSendAPIResponse) {
	v.Reset()
	poolAlibabaAlscChudaTemplateSendAPIResponse.Put(v)
}
