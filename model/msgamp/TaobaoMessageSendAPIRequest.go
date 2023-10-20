package msgamp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomessagesendAPIRequest 消息发送 API请求
// taobao.message.send
//
// 消息发送接口
type TaobaomessagesendAPIRequest struct {
	model.Params
	// 消息发送相关参数
	_sendMessageReq *SendMessageReq
}

// NewTaobaomessagesendRequest 初始化TaobaomessagesendAPIRequest对象
func NewTaobaomessagesendRequest() *TaobaomessagesendAPIRequest {
	return &TaobaomessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendMessageReq is SendMessageReq Setter
// 消息发送相关参数
func (r *TaobaomessagesendAPIRequest) SetSendMessageReq(_sendMessageReq *SendMessageReq) error {
	r._sendMessageReq = _sendMessageReq
	r.Set("send_message_req", _sendMessageReq)
	return nil
}

// GetSendMessageReq SendMessageReq Getter
func (r TaobaomessagesendAPIRequest) GetSendMessageReq() *SendMessageReq {
	return r._sendMessageReq
}
