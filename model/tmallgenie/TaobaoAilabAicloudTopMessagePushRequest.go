package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息中心广播推送消息接口 API请求
taobao.ailab.aicloud.top.message.push

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
*/
type TaobaoAilabAicloudTopMessagePushRequest struct {
    model.Params
    // 消息推送请求
    _messageBroadcastRequest   *MessageBroadcastRequest
    // 当前用户信息
    _userInfoContext   *OpsRequestUserInfoContext
}

// 初始化TaobaoAilabAicloudTopMessagePushRequest对象
func NewTaobaoAilabAicloudTopMessagePushRequest() *TaobaoAilabAicloudTopMessagePushRequest{
    return &TaobaoAilabAicloudTopMessagePushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessagePushRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.push"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MessageBroadcastRequest Setter
// 消息推送请求
func (r *TaobaoAilabAicloudTopMessagePushRequest) SetMessageBroadcastRequest(_messageBroadcastRequest *MessageBroadcastRequest) error {
    r._messageBroadcastRequest = _messageBroadcastRequest
    r.Set("message_broadcast_request", _messageBroadcastRequest)
    return nil
}

// MessageBroadcastRequest Getter
func (r TaobaoAilabAicloudTopMessagePushRequest) GetMessageBroadcastRequest() *MessageBroadcastRequest {
    return r._messageBroadcastRequest
}
// UserInfoContext Setter
// 当前用户信息
func (r *TaobaoAilabAicloudTopMessagePushRequest) SetUserInfoContext(_userInfoContext *OpsRequestUserInfoContext) error {
    r._userInfoContext = _userInfoContext
    r.Set("user_info_context", _userInfoContext)
    return nil
}

// UserInfoContext Getter
func (r TaobaoAilabAicloudTopMessagePushRequest) GetUserInfoContext() *OpsRequestUserInfoContext {
    return r._userInfoContext
}
