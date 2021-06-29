package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处理Usc会员消息接口 APIRequest
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

func NewAlibabaWestcrmMemberMessageHandleRequest() *AlibabaWestcrmMemberMessageHandleRequest{
    return &AlibabaWestcrmMemberMessageHandleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmMemberMessageHandleRequest) GetApiMethodName() string {
    return "alibaba.westcrm.member.message.handle"
}

func (r AlibabaWestcrmMemberMessageHandleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmMemberMessageHandleRequest) SetMessageType(messageType string) error {
    r.messageType = messageType
    r.Set("message_type", messageType)
    return nil
}

func (r AlibabaWestcrmMemberMessageHandleRequest) GetMessageType() string {
    return r.messageType
}

func (r *AlibabaWestcrmMemberMessageHandleRequest) SetMessageContent(messageContent string) error {
    r.messageContent = messageContent
    r.Set("message_content", messageContent)
    return nil
}

func (r AlibabaWestcrmMemberMessageHandleRequest) GetMessageContent() string {
    return r.messageContent
}

