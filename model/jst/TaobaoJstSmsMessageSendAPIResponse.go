package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMessageSendAPIResponse 聚石塔数据paas短信发送接口 API返回值
// taobao.jst.sms.message.send
//
// 聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
type TaobaoJstSmsMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsMessageSendAPIResponseModel).Reset()
}

// TaobaoJstSmsMessageSendAPIResponseModel is 聚石塔数据paas短信发送接口 成功返回结果
type TaobaoJstSmsMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_message_send_response"`
	// 参数错误
	RequestCode string `json:"request_code,omitempty" xml:"request_code,omitempty"`
	// 1234
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 空
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 参数错误
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求成功
	RequestSuccess bool `json:"request_success,omitempty" xml:"request_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstSmsMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RequestCode = ""
	m.RequestId = ""
	m.Module = ""
	m.Message = ""
	m.RequestSuccess = false
}

var poolTaobaoJstSmsMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsMessageSendAPIResponse)
	},
}

// GetTaobaoJstSmsMessageSendAPIResponse 从 sync.Pool 获取 TaobaoJstSmsMessageSendAPIResponse
func GetTaobaoJstSmsMessageSendAPIResponse() *TaobaoJstSmsMessageSendAPIResponse {
	return poolTaobaoJstSmsMessageSendAPIResponse.Get().(*TaobaoJstSmsMessageSendAPIResponse)
}

// ReleaseTaobaoJstSmsMessageSendAPIResponse 将 TaobaoJstSmsMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsMessageSendAPIResponse(v *TaobaoJstSmsMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsMessageSendAPIResponse.Put(v)
}
