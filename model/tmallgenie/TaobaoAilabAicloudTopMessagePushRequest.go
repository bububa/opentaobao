package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息中心广播推送消息接口 APIRequest
taobao.ailab.aicloud.top.message.push

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
*/
type TaobaoAilabAicloudTopMessagePushRequest struct {
    model.Params

    // 消息推送请求
    messageBroadcastRequest   *MessageBroadcastRequest 

    // 当前用户信息
    userInfoContext   *OpsRequestUserInfoContext 

}

func NewTaobaoAilabAicloudTopMessagePushRequest() *TaobaoAilabAicloudTopMessagePushRequest{
    return &TaobaoAilabAicloudTopMessagePushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMessagePushRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.push"
}

func (r TaobaoAilabAicloudTopMessagePushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMessagePushRequest) SetMessageBroadcastRequest(messageBroadcastRequest *MessageBroadcastRequest) error {
    r.messageBroadcastRequest = messageBroadcastRequest
    r.Set("message_broadcast_request", messageBroadcastRequest)
    return nil
}

func (r TaobaoAilabAicloudTopMessagePushRequest) GetMessageBroadcastRequest() *MessageBroadcastRequest {
    return r.messageBroadcastRequest
}

func (r *TaobaoAilabAicloudTopMessagePushRequest) SetUserInfoContext(userInfoContext *OpsRequestUserInfoContext) error {
    r.userInfoContext = userInfoContext
    r.Set("user_info_context", userInfoContext)
    return nil
}

func (r TaobaoAilabAicloudTopMessagePushRequest) GetUserInfoContext() *OpsRequestUserInfoContext {
    return r.userInfoContext
}

