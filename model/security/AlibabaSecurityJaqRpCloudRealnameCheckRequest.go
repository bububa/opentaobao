package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
验证姓名和证件号 APIRequest
alibaba.security.jaq.rp.cloud.realname.check

验证姓名和证件号
*/
type AlibabaSecurityJaqRpCloudRealnameCheckRequest struct {
    model.Params

    // token
    verifyToken   string 

    // 要识别的信息
    imageUrls   string 

    // 姓名
    name   string 

    // 证件号
    identityCode   string 

}

func NewAlibabaSecurityJaqRpCloudRealnameCheckRequest() *AlibabaSecurityJaqRpCloudRealnameCheckRequest{
    return &AlibabaSecurityJaqRpCloudRealnameCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.realname.check"
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetImageUrls() string {
    return r.imageUrls
}

func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetName() string {
    return r.name
}

func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetIdentityCode(identityCode string) error {
    r.identityCode = identityCode
    r.Set("identity_code", identityCode)
    return nil
}

func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetIdentityCode() string {
    return r.identityCode
}

