package tmallgenie

import (
	"net/url"

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
	// 当前用户信息
	_userInfoContext *OpsRequestUserInfoContext
}

// NewTaobaoAilabAicloudTopMessagePushRequest 初始化TaobaoAilabAicloudTopMessagePushAPIRequest对象
func NewTaobaoAilabAicloudTopMessagePushRequest() *TaobaoAilabAicloudTopMessagePushAPIRequest {
	return &TaobaoAilabAicloudTopMessagePushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

// SetUserInfoContext is UserInfoContext Setter
// 当前用户信息
func (r *TaobaoAilabAicloudTopMessagePushAPIRequest) SetUserInfoContext(_userInfoContext *OpsRequestUserInfoContext) error {
	r._userInfoContext = _userInfoContext
	r.Set("user_info_context", _userInfoContext)
	return nil
}

// GetUserInfoContext UserInfoContext Getter
func (r TaobaoAilabAicloudTopMessagePushAPIRequest) GetUserInfoContext() *OpsRequestUserInfoContext {
	return r._userInfoContext
}
