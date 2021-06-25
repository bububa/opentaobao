package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询认证结果 APIRequest
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果
*/
type AlibabaSecurityJaqRpQueryRequest struct {
    model.Params

    // token
    verifyToken   string 

}

func NewAlibabaSecurityJaqRpQueryRequest() *AlibabaSecurityJaqRpQueryRequest{
    return &AlibabaSecurityJaqRpQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpQueryRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.query"
}

func (r AlibabaSecurityJaqRpQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpQueryRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpQueryRequest) GetVerifyToken() string {
    return r.verifyToken
}

