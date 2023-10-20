package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagepushAPIRequest 天猫精灵消息中心广播推送消息接口 API请求
// taobao.ailab.aicloud.top.message.push
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
type TaobaoailabaicloudtopmessagepushAPIRequest struct {
	model.Params
	// 消息推送请求
	_messageBroadcastRequest *MessageBroadcastRequest
}

// NewTaobaoailabaicloudtopmessagepushRequest 初始化TaobaoailabaicloudtopmessagepushAPIRequest对象
func NewTaobaoailabaicloudtopmessagepushRequest() *TaobaoailabaicloudtopmessagepushAPIRequest {
	return &TaobaoailabaicloudtopmessagepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmessagepushAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmessagepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmessagepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageBroadcastRequest is MessageBroadcastRequest Setter
// 消息推送请求
func (r *TaobaoailabaicloudtopmessagepushAPIRequest) SetMessageBroadcastRequest(_messageBroadcastRequest *MessageBroadcastRequest) error {
	r._messageBroadcastRequest = _messageBroadcastRequest
	r.Set("message_broadcast_request", _messageBroadcastRequest)
	return nil
}

// GetMessageBroadcastRequest MessageBroadcastRequest Getter
func (r TaobaoailabaicloudtopmessagepushAPIRequest) GetMessageBroadcastRequest() *MessageBroadcastRequest {
	return r._messageBroadcastRequest
}
