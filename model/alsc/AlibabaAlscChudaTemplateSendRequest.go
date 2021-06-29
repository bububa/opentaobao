package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活触达模板消息发送接口 APIRequest
alibaba.alsc.chuda.template.send

允许三方小程序通过该api发送本地生活触达消息
*/
type AlibabaAlscChudaTemplateSendRequest struct {
    model.Params

    // 请求参数
    notifyRequest   *TemplateNotifyRequest 

}

func NewAlibabaAlscChudaTemplateSendRequest() *AlibabaAlscChudaTemplateSendRequest{
    return &AlibabaAlscChudaTemplateSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscChudaTemplateSendRequest) GetApiMethodName() string {
    return "alibaba.alsc.chuda.template.send"
}

func (r AlibabaAlscChudaTemplateSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscChudaTemplateSendRequest) SetNotifyRequest(notifyRequest *TemplateNotifyRequest) error {
    r.notifyRequest = notifyRequest
    r.Set("notify_request", notifyRequest)
    return nil
}

func (r AlibabaAlscChudaTemplateSendRequest) GetNotifyRequest() *TemplateNotifyRequest {
    return r.notifyRequest
}

