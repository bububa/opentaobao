package msgamp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息发送 APIRequest
taobao.message.send

消息发送接口
*/
type TaobaoMessageSendRequest struct {
    model.Params

    // 消息发送相关参数
    sendMessageReq   *SendMessageReq 

}

func NewTaobaoMessageSendRequest() *TaobaoMessageSendRequest{
    return &TaobaoMessageSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMessageSendRequest) GetApiMethodName() string {
    return "taobao.message.send"
}

func (r TaobaoMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMessageSendRequest) SetSendMessageReq(sendMessageReq *SendMessageReq) error {
    r.sendMessageReq = sendMessageReq
    r.Set("send_message_req", sendMessageReq)
    return nil
}

func (r TaobaoMessageSendRequest) GetSendMessageReq() *SendMessageReq {
    return r.sendMessageReq
}

