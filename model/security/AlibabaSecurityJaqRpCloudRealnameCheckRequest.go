package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
验证姓名和证件号 API请求
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

// 初始化AlibabaSecurityJaqRpCloudRealnameCheckRequest对象
func NewAlibabaSecurityJaqRpCloudRealnameCheckRequest() *AlibabaSecurityJaqRpCloudRealnameCheckRequest{
    return &AlibabaSecurityJaqRpCloudRealnameCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.realname.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetVerifyToken() string {
    return r.verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetImageUrls() string {
    return r.imageUrls
}
// Name Setter
// 姓名
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetName() string {
    return r.name
}
// IdentityCode Setter
// 证件号
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetIdentityCode(identityCode string) error {
    r.identityCode = identityCode
    r.Set("identity_code", identityCode)
    return nil
}

// IdentityCode Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetIdentityCode() string {
    return r.identityCode
}
