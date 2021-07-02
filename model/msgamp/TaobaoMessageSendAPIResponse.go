package msgamp

import (
	"encoding/xml"

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

// TaobaoMessageSendAPIResponseModel is 消息发送 成功返回结果
type TaobaoMessageSendAPIResponseModel struct {
	XMLName xml.Name `xml:"message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
