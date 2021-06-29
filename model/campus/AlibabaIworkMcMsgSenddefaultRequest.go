package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
给注册用户发送消息 API请求
alibaba.iwork.mc.msg.senddefault

给神鲸注册用户发送对应操作结果的消息
*/
type AlibabaIworkMcMsgSenddefaultRequest struct {
    model.Params
    // 消息对象
    _messageEvent   *DefaultMessageEvent
}

// 初始化AlibabaIworkMcMsgSenddefaultRequest对象
func NewAlibabaIworkMcMsgSenddefaultRequest() *AlibabaIworkMcMsgSenddefaultRequest{
    return &AlibabaIworkMcMsgSenddefaultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIworkMcMsgSenddefaultRequest) GetApiMethodName() string {
    return "alibaba.iwork.mc.msg.senddefault"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIworkMcMsgSenddefaultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MessageEvent Setter
// 消息对象
func (r *AlibabaIworkMcMsgSenddefaultRequest) SetMessageEvent(_messageEvent *DefaultMessageEvent) error {
    r._messageEvent = _messageEvent
    r.Set("message_event", _messageEvent)
    return nil
}

// MessageEvent Getter
func (r AlibabaIworkMcMsgSenddefaultRequest) GetMessageEvent() *DefaultMessageEvent {
    return r._messageEvent
}
