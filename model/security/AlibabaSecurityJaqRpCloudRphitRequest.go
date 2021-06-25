package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务日志打点 APIRequest
alibaba.security.jaq.rp.cloud.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpCloudRphitRequest struct {
    model.Params

    // xxx
    content   string 

}

func NewAlibabaSecurityJaqRpCloudRphitRequest() *AlibabaSecurityJaqRpCloudRphitRequest{
    return &AlibabaSecurityJaqRpCloudRphitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudRphitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.rphit"
}

func (r AlibabaSecurityJaqRpCloudRphitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudRphitRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaSecurityJaqRpCloudRphitRequest) GetContent() string {
    return r.content
}

