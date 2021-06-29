package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息给手机用户 API请求
alibaba.iwork.mc.msg.sendmobile

给手机用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSendmobileRequest struct {
    model.Params
    // 消息对象
    mobileReceiverMessageEvent   *MobileReceiverMessageEvent
}

// 初始化AlibabaIworkMcMsgSendmobileRequest对象
func NewAlibabaIworkMcMsgSendmobileRequest() *AlibabaIworkMcMsgSendmobileRequest{
    return &AlibabaIworkMcMsgSendmobileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIworkMcMsgSendmobileRequest) GetApiMethodName() string {
    return "alibaba.iwork.mc.msg.sendmobile"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIworkMcMsgSendmobileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MobileReceiverMessageEvent Setter
// 消息对象
func (r *AlibabaIworkMcMsgSendmobileRequest) SetMobileReceiverMessageEvent(mobileReceiverMessageEvent *MobileReceiverMessageEvent) error {
    r.mobileReceiverMessageEvent = mobileReceiverMessageEvent
    r.Set("mobile_receiver_message_event", mobileReceiverMessageEvent)
    return nil
}

// MobileReceiverMessageEvent Getter
func (r AlibabaIworkMcMsgSendmobileRequest) GetMobileReceiverMessageEvent() *MobileReceiverMessageEvent {
    return r.mobileReceiverMessageEvent
}
