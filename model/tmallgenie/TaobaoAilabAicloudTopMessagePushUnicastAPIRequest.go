package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessagePushUnicastAPIRequest 天猫精灵消息中心单播推送消息接口 API请求
// taobao.ailab.aicloud.top.message.push.unicast
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
type TaobaoAilabAicloudTopMessagePushUnicastAPIRequest struct {
	model.Params
	// 消息推送单播请求体
	_messageUnicastRequest *MessageUnicastRequest
}

// NewTaobaoAilabAicloudTopMessagePushUnicastRequest 初始化TaobaoAilabAicloudTopMessagePushUnicastAPIRequest对象
func NewTaobaoAilabAicloudTopMessagePushUnicastRequest() *TaobaoAilabAicloudTopMessagePushUnicastAPIRequest {
	return &TaobaoAilabAicloudTopMessagePushUnicastAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push.unicast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMessageUnicastRequest is MessageUnicastRequest Setter
// 消息推送单播请求体
func (r *TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) SetMessageUnicastRequest(_messageUnicastRequest *MessageUnicastRequest) error {
	r._messageUnicastRequest = _messageUnicastRequest
	r.Set("message_unicast_request", _messageUnicastRequest)
	return nil
}

// GetMessageUnicastRequest MessageUnicastRequest Getter
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetMessageUnicastRequest() *MessageUnicastRequest {
	return r._messageUnicastRequest
}
