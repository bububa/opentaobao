package msgamp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMessageSendAPIResponse 消息发送 API返回值
// taobao.message.send
//
// 消息发送接口
type TaobaoMessageSendAPIResponse struct {
	model.CommonResponse
	TaobaoMessageSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMessageSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMessageSendAPIResponseModel).Reset()
}

// TaobaoMessageSendAPIResponseModel is 消息发送 成功返回结果
type TaobaoMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMessageSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.IsSuccess = false
}

var poolTaobaoMessageSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMessageSendAPIResponse)
	},
}

// GetTaobaoMessageSendAPIResponse 从 sync.Pool 获取 TaobaoMessageSendAPIResponse
func GetTaobaoMessageSendAPIResponse() *TaobaoMessageSendAPIResponse {
	return poolTaobaoMessageSendAPIResponse.Get().(*TaobaoMessageSendAPIResponse)
}

// ReleaseTaobaoMessageSendAPIResponse 将 TaobaoMessageSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMessageSendAPIResponse(v *TaobaoMessageSendAPIResponse) {
	v.Reset()
	poolTaobaoMessageSendAPIResponse.Put(v)
}
