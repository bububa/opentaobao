package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsOaidMessageSendAPIResponse 基于OAID的短信发送接口 API返回值
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
type TaobaoJstSmsOaidMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsOaidMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsOaidMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsOaidMessageSendAPIResponseModel).Reset()
}

// TaobaoJstSmsOaidMessageSendAPIResponseModel is 基于OAID的短信发送接口 成功返回结果
type TaobaoJstSmsOaidMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_oaid_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信拓展码
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// top请求id
	ReqId string `json:"req_id,omitempty" xml:"req_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstSmsOaidMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = ""
	m.ReqId = ""
}

var poolTaobaoJstSmsOaidMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsOaidMessageSendAPIResponse)
	},
}

// GetTaobaoJstSmsOaidMessageSendAPIResponse 从 sync.Pool 获取 TaobaoJstSmsOaidMessageSendAPIResponse
func GetTaobaoJstSmsOaidMessageSendAPIResponse() *TaobaoJstSmsOaidMessageSendAPIResponse {
	return poolTaobaoJstSmsOaidMessageSendAPIResponse.Get().(*TaobaoJstSmsOaidMessageSendAPIResponse)
}

// ReleaseTaobaoJstSmsOaidMessageSendAPIResponse 将 TaobaoJstSmsOaidMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsOaidMessageSendAPIResponse(v *TaobaoJstSmsOaidMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsOaidMessageSendAPIResponse.Put(v)
}
