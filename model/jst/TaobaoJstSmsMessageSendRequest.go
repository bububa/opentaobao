package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔数据paas短信发送接口 API请求
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。
*/
type TaobaoJstSmsMessageSendRequest struct {
    model.Params
    // 短信发送请求
    _sendMessageRequest   *SendMessageRequest
}

// 初始化TaobaoJstSmsMessageSendRequest对象
func NewTaobaoJstSmsMessageSendRequest() *TaobaoJstSmsMessageSendRequest{
    return &TaobaoJstSmsMessageSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendMessageRequest Setter
// 短信发送请求
func (r *TaobaoJstSmsMessageSendRequest) SetSendMessageRequest(_sendMessageRequest *SendMessageRequest) error {
    r._sendMessageRequest = _sendMessageRequest
    r.Set("send_message_request", _sendMessageRequest)
    return nil
}

// SendMessageRequest Getter
func (r TaobaoJstSmsMessageSendRequest) GetSendMessageRequest() *SendMessageRequest {
    return r._sendMessageRequest
}
