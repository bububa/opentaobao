package alimsg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虚拟发货消息发送接口 APIRequest
alibaba.idle.order.msg.send

用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
*/
type AlibabaIdleOrderMsgSendRequest struct {
    model.Params

    // 订单id
    orderId   string 

    // 消息发送内容
    text   string 

}

func NewAlibabaIdleOrderMsgSendRequest() *AlibabaIdleOrderMsgSendRequest{
    return &AlibabaIdleOrderMsgSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleOrderMsgSendRequest) GetApiMethodName() string {
    return "alibaba.idle.order.msg.send"
}

func (r AlibabaIdleOrderMsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleOrderMsgSendRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaIdleOrderMsgSendRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaIdleOrderMsgSendRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r AlibabaIdleOrderMsgSendRequest) GetText() string {
    return r.text
}

