package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云上传接口 APIRequest
alibaba.security.jaq.rp.cloud.upload

聚安全实人认证上传认证信息
*/
type AlibabaSecurityJaqRpCloudUploadRequest struct {
    model.Params

    // 认证token
    verifyToken   string 

    // []
    elements   []Elements 

}

func NewAlibabaSecurityJaqRpCloudUploadRequest() *AlibabaSecurityJaqRpCloudUploadRequest{
    return &AlibabaSecurityJaqRpCloudUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudUploadRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.upload"
}

func (r AlibabaSecurityJaqRpCloudUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudUploadRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudUploadRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpCloudUploadRequest) SetElements(elements []Elements) error {
    r.elements = elements
    r.Set("elements", elements)
    return nil
}

func (r AlibabaSecurityJaqRpCloudUploadRequest) GetElements() []Elements {
    return r.elements
}

