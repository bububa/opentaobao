package msgamp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMessageSendAPIRequest
消息发送 API请求
taobao.message.send

消息发送接口 */
type TaobaoMessageSendAPIRequest struct {
	model.Params
	// 消息发送相关参数
	_sendMessageReq *SendMessageReq
}

// NewTaobaoMessageSendRequest 初始化TaobaoMessageSendAPIRequest对象
func NewTaobaoMessageSendRequest() *TaobaoMessageSendAPIRequest {
	return &TaobaoMessageSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMessageSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SendMessageReq Setter
// 消息发送相关参数
func (r *TaobaoMessageSendAPIRequest) SetSendMessageReq(_sendMessageReq *SendMessageReq) error {
	r._sendMessageReq = _sendMessageReq
	r.Set("send_message_req", _sendMessageReq)
	return nil
}

// Get SendMessageReq Getter
func (r TaobaoMessageSendAPIRequest) GetSendMessageReq() *SendMessageReq {
	return r._sendMessageReq
}
