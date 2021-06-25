package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证上传认证信息 APIRequest
alibaba.security.jaq.rp.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpUploadRequest struct {
    model.Params

    // 认证会话token
    verifyToken   string 

    // 此次需要上传的认证信息的列表
    elements   []Element 

}

func NewAlibabaSecurityJaqRpUploadRequest() *AlibabaSecurityJaqRpUploadRequest{
    return &AlibabaSecurityJaqRpUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpUploadRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.upload"
}

func (r AlibabaSecurityJaqRpUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpUploadRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpUploadRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpUploadRequest) SetElements(elements []Element) error {
    r.elements = elements
    r.Set("elements", elements)
    return nil
}

func (r AlibabaSecurityJaqRpUploadRequest) GetElements() []Element {
    return r.elements
}

