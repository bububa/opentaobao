package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全-实人认证日志打点接口 API请求
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpRphitRequest struct {
    model.Params
    // xxx
    _content   string
}

// 初始化AlibabaSecurityJaqRpRphitRequest对象
func NewAlibabaSecurityJaqRpRphitRequest() *AlibabaSecurityJaqRpRphitRequest{
    return &AlibabaSecurityJaqRpRphitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpRphitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.rphit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpRphitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// xxx
func (r *AlibabaSecurityJaqRpRphitRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaSecurityJaqRpRphitRequest) GetContent() string {
    return r._content
}
