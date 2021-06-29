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
    _verifyToken   string
    // 要识别的信息
    _imageUrls   string
    // 姓名
    _name   string
    // 证件号
    _identityCode   string
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
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetVerifyToken() string {
    return r._verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetImageUrls(_imageUrls string) error {
    r._imageUrls = _imageUrls
    r.Set("image_urls", _imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetImageUrls() string {
    return r._imageUrls
}
// Name Setter
// 姓名
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetName() string {
    return r._name
}
// IdentityCode Setter
// 证件号
func (r *AlibabaSecurityJaqRpCloudRealnameCheckRequest) SetIdentityCode(_identityCode string) error {
    r._identityCode = _identityCode
    r.Set("identity_code", _identityCode)
    return nil
}

// IdentityCode Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckRequest) GetIdentityCode() string {
    return r._identityCode
}
