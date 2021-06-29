package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
本地生活触达模板消息发送接口 API请求
alibaba.alsc.chuda.template.send

允许三方小程序通过该api发送本地生活触达消息
*/
type AlibabaAlscChudaTemplateSendRequest struct {
    model.Params
    // 请求参数
    _notifyRequest   *TemplateNotifyRequest
}

// 初始化AlibabaAlscChudaTemplateSendRequest对象
func NewAlibabaAlscChudaTemplateSendRequest() *AlibabaAlscChudaTemplateSendRequest{
    return &AlibabaAlscChudaTemplateSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscChudaTemplateSendRequest) GetApiMethodName() string {
    return "alibaba.alsc.chuda.template.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscChudaTemplateSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NotifyRequest Setter
// 请求参数
func (r *AlibabaAlscChudaTemplateSendRequest) SetNotifyRequest(_notifyRequest *TemplateNotifyRequest) error {
    r._notifyRequest = _notifyRequest
    r.Set("notify_request", _notifyRequest)
    return nil
}

// NotifyRequest Getter
func (r AlibabaAlscChudaTemplateSendRequest) GetNotifyRequest() *TemplateNotifyRequest {
    return r._notifyRequest
}
