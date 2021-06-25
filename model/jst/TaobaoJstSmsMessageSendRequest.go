package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔数据paas短信发送接口 APIRequest
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
*/
type TaobaoJstSmsMessageSendRequest struct {
    model.Params

    // 短信发送请求
    sendMessageRequest   *SendMessageRequest 

}

func NewTaobaoJstSmsMessageSendRequest() *TaobaoJstSmsMessageSendRequest{
    return &TaobaoJstSmsMessageSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.send"
}

func (r TaobaoJstSmsMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsMessageSendRequest) SetSendMessageRequest(sendMessageRequest *SendMessageRequest) error {
    r.sendMessageRequest = sendMessageRequest
    r.Set("send_message_request", sendMessageRequest)
    return nil
}

func (r TaobaoJstSmsMessageSendRequest) GetSendMessageRequest() *SendMessageRequest {
    return r.sendMessageRequest
}

