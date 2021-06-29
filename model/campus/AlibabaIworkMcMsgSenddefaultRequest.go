package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给注册用户发送消息 APIRequest
alibaba.iwork.mc.msg.senddefault

给神鲸注册用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSenddefaultRequest struct {
    model.Params

    // 消息对象
    messageEvent   *DefaultMessageEvent 

}

func NewAlibabaIworkMcMsgSenddefaultRequest() *AlibabaIworkMcMsgSenddefaultRequest{
    return &AlibabaIworkMcMsgSenddefaultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIworkMcMsgSenddefaultRequest) GetApiMethodName() string {
    return "alibaba.iwork.mc.msg.senddefault"
}

func (r AlibabaIworkMcMsgSenddefaultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIworkMcMsgSenddefaultRequest) SetMessageEvent(messageEvent *DefaultMessageEvent) error {
    r.messageEvent = messageEvent
    r.Set("message_event", messageEvent)
    return nil
}

func (r AlibabaIworkMcMsgSenddefaultRequest) GetMessageEvent() *DefaultMessageEvent {
    return r.messageEvent
}

