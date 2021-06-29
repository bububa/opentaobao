package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处理Usc会员消息接口 API请求
alibaba.westcrm.member.message.handle

处理Usc会员消息接口
*/
type AlibabaWestcrmMemberMessageHandleRequest struct {
    model.Params
    // 消息类型
    messageType   string
    // 消息内容
    messageContent   string
}

// 初始化AlibabaWestcrmMemberMessageHandleRequest对象
func NewAlibabaWestcrmMemberMessageHandleRequest() *AlibabaWestcrmMemberMessageHandleRequest{
    return &AlibabaWestcrmMemberMessageHandleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMemberMessageHandleRequest) GetApiMethodName() string {
    return "alibaba.westcrm.member.message.handle"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMemberMessageHandleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MessageType Setter
// 消息类型
func (r *AlibabaWestcrmMemberMessageHandleRequest) SetMessageType(messageType string) error {
    r.messageType = messageType
    r.Set("message_type", messageType)
    return nil
}

// MessageType Getter
func (r AlibabaWestcrmMemberMessageHandleRequest) GetMessageType() string {
    return r.messageType
}
// MessageContent Setter
// 消息内容
func (r *AlibabaWestcrmMemberMessageHandleRequest) SetMessageContent(messageContent string) error {
    r.messageContent = messageContent
    r.Set("message_content", messageContent)
    return nil
}

// MessageContent Getter
func (r AlibabaWestcrmMemberMessageHandleRequest) GetMessageContent() string {
    return r.messageContent
}
