package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息中心单播推送消息接口 APIRequest
taobao.ailab.aicloud.top.message.push.unicast

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
*/
type TaobaoAilabAicloudTopMessagePushUnicastRequest struct {
    model.Params

    // 消息推送单播请求体
    messageUnicastRequest   *MessageUnicastRequest 

}

func NewTaobaoAilabAicloudTopMessagePushUnicastRequest() *TaobaoAilabAicloudTopMessagePushUnicastRequest{
    return &TaobaoAilabAicloudTopMessagePushUnicastRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMessagePushUnicastRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.push.unicast"
}

func (r TaobaoAilabAicloudTopMessagePushUnicastRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMessagePushUnicastRequest) SetMessageUnicastRequest(messageUnicastRequest *MessageUnicastRequest) error {
    r.messageUnicastRequest = messageUnicastRequest
    r.Set("message_unicast_request", messageUnicastRequest)
    return nil
}

func (r TaobaoAilabAicloudTopMessagePushUnicastRequest) GetMessageUnicastRequest() *MessageUnicastRequest {
    return r.messageUnicastRequest
}

