package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest 验证姓名和证件号 API请求
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
type AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 姓名
	_name string
	// 证件号
	_identityCode string
	// 要识别的信息
	_imageUrls string
}

// NewAlibabaSecurityJaqRpCloudRealnameCheckRequest 初始化AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest对象
func NewAlibabaSecurityJaqRpCloudRealnameCheckRequest() *AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest {
	return &AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.realname.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetName is Name Setter
// 姓名
func (r *AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetName() string {
	return r._name
}

// SetIdentityCode is IdentityCode Setter
// 证件号
func (r *AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) SetIdentityCode(_identityCode string) error {
	r._identityCode = _identityCode
	r.Set("identity_code", _identityCode)
	return nil
}

// GetIdentityCode IdentityCode Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetIdentityCode() string {
	return r._identityCode
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest) GetImageUrls() string {
	return r._imageUrls
}
