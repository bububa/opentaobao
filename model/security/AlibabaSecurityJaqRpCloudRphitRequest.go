package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务日志打点 API请求
alibaba.security.jaq.rp.cloud.rphit

聚安全实人认证日志打点接口
*/
type AlibabaSecurityJaqRpCloudRphitRequest struct {
    model.Params
    // xxx
    content   string
}

// 初始化AlibabaSecurityJaqRpCloudRphitRequest对象
func NewAlibabaSecurityJaqRpCloudRphitRequest() *AlibabaSecurityJaqRpCloudRphitRequest{
    return &AlibabaSecurityJaqRpCloudRphitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudRphitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.rphit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudRphitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// xxx
func (r *AlibabaSecurityJaqRpCloudRphitRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaSecurityJaqRpCloudRphitRequest) GetContent() string {
    return r.content
}
