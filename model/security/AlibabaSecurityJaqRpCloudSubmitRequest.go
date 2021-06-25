package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务提交接口 APIRequest
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpCloudSubmitRequest struct {
    model.Params

    // 认证token
    verifyToken   string 

}

func NewAlibabaSecurityJaqRpCloudSubmitRequest() *AlibabaSecurityJaqRpCloudSubmitRequest{
    return &AlibabaSecurityJaqRpCloudSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.submit"
}

func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudSubmitRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetVerifyToken() string {
    return r.verifyToken
}

