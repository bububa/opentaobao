package msgamp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息发送 API请求
taobao.message.send

消息发送接口
*/
type TaobaoMessageSendRequest struct {
    model.Params
    // 消息发送相关参数
    _sendMessageReq   *SendMessageReq
}

// 初始化TaobaoMessageSendRequest对象
func NewTaobaoMessageSendRequest() *TaobaoMessageSendRequest{
    return &TaobaoMessageSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMessageSendRequest) GetApiMethodName() string {
    return "taobao.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendMessageReq Setter
// 消息发送相关参数
func (r *TaobaoMessageSendRequest) SetSendMessageReq(_sendMessageReq *SendMessageReq) error {
    r._sendMessageReq = _sendMessageReq
    r.Set("send_message_req", _sendMessageReq)
    return nil
}

// SendMessageReq Getter
func (r TaobaoMessageSendRequest) GetSendMessageReq() *SendMessageReq {
    return r._sendMessageReq
}
