package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessagePushAPIRequest 天猫精灵消息中心广播推送消息接口 API请求
// taobao.ailab.aicloud.top.message.push
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
type TaobaoAilabAicloudTopMessagePushAPIRequest struct {
	model.Params
	// 消息推送请求
	_messageBroadcastRequest *MessageBroadcastRequest
}

// NewTaobaoAilabAicloudTopMessagePushRequest 初始化TaobaoAilabAicloudTopMessagePushAPIRequest对象
func NewTaobaoAilabAicloudTopMessagePushRequest() *TaobaoAilabAicloudTopMessagePushAPIRequest {
	return &TaobaoAilabAicloudTopMessagePushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMessagePushAPIRequest) Reset() {
	r._messageBroadcastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageBroadcastRequest is MessageBroadcastRequest Setter
// 消息推送请求
func (r *TaobaoAilabAicloudTopMessagePushAPIRequest) SetMessageBroadcastRequest(_messageBroadcastRequest *MessageBroadcastRequest) error {
	r._messageBroadcastRequest = _messageBroadcastRequest
	r.Set("message_broadcast_request", _messageBroadcastRequest)
	return nil
}

// GetMessageBroadcastRequest MessageBroadcastRequest Getter
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetMessageBroadcastRequest() *MessageBroadcastRequest {
	return r._messageBroadcastRequest
}

var poolTaobaoAilabAicloudTopMessagePushAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMessagePushRequest()
	},
}

// GetTaobaoAilabAicloudTopMessagePushRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMessagePushAPIRequest
func GetTaobaoAilabAicloudTopMessagePushAPIRequest() *TaobaoAilabAicloudTopMessagePushAPIRequest {
	return poolTaobaoAilabAicloudTopMessagePushAPIRequest.Get().(*TaobaoAilabAicloudTopMessagePushAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMessagePushAPIRequest 将 TaobaoAilabAicloudTopMessagePushAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessagePushAPIRequest(v *TaobaoAilabAicloudTopMessagePushAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessagePushAPIRequest.Put(v)
}
