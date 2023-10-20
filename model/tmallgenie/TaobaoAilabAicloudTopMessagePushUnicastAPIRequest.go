package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagepushunicastAPIRequest 天猫精灵消息中心单播推送消息接口 API请求
// taobao.ailab.aicloud.top.message.push.unicast
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
type TaobaoailabaicloudtopmessagepushunicastAPIRequest struct {
	model.Params
	// 消息推送单播请求体
	_messageUnicastRequest *MessageUnicastRequest
}

// NewTaobaoailabaicloudtopmessagepushunicastRequest 初始化TaobaoailabaicloudtopmessagepushunicastAPIRequest对象
func NewTaobaoailabaicloudtopmessagepushunicastRequest() *TaobaoailabaicloudtopmessagepushunicastAPIRequest {
	return &TaobaoailabaicloudtopmessagepushunicastAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmessagepushunicastAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push.unicast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmessagepushunicastAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmessagepushunicastAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessageUnicastRequest is MessageUnicastRequest Setter
// 消息推送单播请求体
func (r *TaobaoailabaicloudtopmessagepushunicastAPIRequest) SetMessageUnicastRequest(_messageUnicastRequest *MessageUnicastRequest) error {
	r._messageUnicastRequest = _messageUnicastRequest
	r.Set("message_unicast_request", _messageUnicastRequest)
	return nil
}

// GetMessageUnicastRequest MessageUnicastRequest Getter
func (r TaobaoailabaicloudtopmessagepushunicastAPIRequest) GetMessageUnicastRequest() *MessageUnicastRequest {
	return r._messageUnicastRequest
}
