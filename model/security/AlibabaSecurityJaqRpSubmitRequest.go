package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证提交认证接口 APIRequest
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpSubmitRequest struct {
    model.Params

    // 认证token
    verifyToken   string 

}

func NewAlibabaSecurityJaqRpSubmitRequest() *AlibabaSecurityJaqRpSubmitRequest{
    return &AlibabaSecurityJaqRpSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpSubmitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.submit"
}

func (r AlibabaSecurityJaqRpSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpSubmitRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpSubmitRequest) GetVerifyToken() string {
    return r.verifyToken
}

