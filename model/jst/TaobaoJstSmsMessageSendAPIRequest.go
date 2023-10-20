package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsmessagesendAPIRequest 聚石塔数据paas短信发送接口 API请求
// taobao.jst.sms.message.send
//
// 聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
type TaobaojstsmsmessagesendAPIRequest struct {
	model.Params
	// 短信发送请求
	_sendMessageRequest *SendMessageRequest
}

// NewTaobaojstsmsmessagesendRequest 初始化TaobaojstsmsmessagesendAPIRequest对象
func NewTaobaojstsmsmessagesendRequest() *TaobaojstsmsmessagesendAPIRequest {
	return &TaobaojstsmsmessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmsmessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmsmessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmsmessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendMessageRequest is SendMessageRequest Setter
// 短信发送请求
func (r *TaobaojstsmsmessagesendAPIRequest) SetSendMessageRequest(_sendMessageRequest *SendMessageRequest) error {
	r._sendMessageRequest = _sendMessageRequest
	r.Set("send_message_request", _sendMessageRequest)
	return nil
}

// GetSendMessageRequest SendMessageRequest Getter
func (r TaobaojstsmsmessagesendAPIRequest) GetSendMessageRequest() *SendMessageRequest {
	return r._sendMessageRequest
}
