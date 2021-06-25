package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全-实人认证日志打点接口 APIRequest
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpRphitRequest struct {
    model.Params

    // xxx
    content   string 

}

func NewAlibabaSecurityJaqRpRphitRequest() *AlibabaSecurityJaqRpRphitRequest{
    return &AlibabaSecurityJaqRpRphitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpRphitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.rphit"
}

func (r AlibabaSecurityJaqRpRphitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpRphitRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaSecurityJaqRpRphitRequest) GetContent() string {
    return r.content
}

