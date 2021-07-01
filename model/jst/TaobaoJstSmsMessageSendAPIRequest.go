package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageSendAPIRequest
聚石塔数据paas短信发送接口 API请求
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。 */
type TaobaoJstSmsMessageSendAPIRequest struct {
	model.Params
	// 短信发送请求
	_sendMessageRequest *SendMessageRequest
}

// NewTaobaoJstSmsMessageSendRequest 初始化TaobaoJstSmsMessageSendAPIRequest对象
func NewTaobaoJstSmsMessageSendRequest() *TaobaoJstSmsMessageSendAPIRequest {
	return &TaobaoJstSmsMessageSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SendMessageRequest Setter
// 短信发送请求
func (r *TaobaoJstSmsMessageSendAPIRequest) SetSendMessageRequest(_sendMessageRequest *SendMessageRequest) error {
	r._sendMessageRequest = _sendMessageRequest
	r.Set("send_message_request", _sendMessageRequest)
	return nil
}

// Get SendMessageRequest Getter
func (r TaobaoJstSmsMessageSendAPIRequest) GetSendMessageRequest() *SendMessageRequest {
	return r._sendMessageRequest
}
