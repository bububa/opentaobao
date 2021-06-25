package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取结果接口 APIRequest
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
type AlibabaSecurityJaqRpFetchmaterialRequest struct {
    model.Params

    // 消息服务推送的key
    securityKey   string 

}

func NewAlibabaSecurityJaqRpFetchmaterialRequest() *AlibabaSecurityJaqRpFetchmaterialRequest{
    return &AlibabaSecurityJaqRpFetchmaterialRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.fetchmaterial"
}

func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpFetchmaterialRequest) SetSecurityKey(securityKey string) error {
    r.securityKey = securityKey
    r.Set("security_key", securityKey)
    return nil
}

func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetSecurityKey() string {
    return r.securityKey
}

