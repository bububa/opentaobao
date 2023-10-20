package tmallgenie

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) Reset() {
	r._messageUnicastRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push.unicast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoAilabAicloudTopMessagePushUnicastAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMessagePushUnicastRequest()
	},
}

// GetTaobaoAilabAicloudTopMessagePushUnicastRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMessagePushUnicastAPIRequest
func GetTaobaoAilabAicloudTopMessagePushUnicastAPIRequest() *TaobaoAilabAicloudTopMessagePushUnicastAPIRequest {
	return poolTaobaoAilabAicloudTopMessagePushUnicastAPIRequest.Get().(*TaobaoAilabAicloudTopMessagePushUnicastAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMessagePushUnicastAPIRequest 将 TaobaoAilabAicloudTopMessagePushUnicastAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessagePushUnicastAPIRequest(v *TaobaoAilabAicloudTopMessagePushUnicastAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessagePushUnicastAPIRequest.Put(v)
}
