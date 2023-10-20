package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudrealnamecheckAPIRequest 验证姓名和证件号 API请求
// alibaba.security.jaq.rp.cloud.realname.check
//
// 验证姓名和证件号
type AlibabasecurityjaqrpcloudrealnamecheckAPIRequest struct {
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

// NewAlibabasecurityjaqrpcloudrealnamecheckRequest 初始化AlibabasecurityjaqrpcloudrealnamecheckAPIRequest对象
func NewAlibabasecurityjaqrpcloudrealnamecheckRequest() *AlibabasecurityjaqrpcloudrealnamecheckAPIRequest {
	return &AlibabasecurityjaqrpcloudrealnamecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.realname.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetName is Name Setter
// 姓名
func (r *AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetName() string {
	return r._name
}

// SetIdentityCode is IdentityCode Setter
// 证件号
func (r *AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) SetIdentityCode(_identityCode string) error {
	r._identityCode = _identityCode
	r.Set("identity_code", _identityCode)
	return nil
}

// GetIdentityCode IdentityCode Getter
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetIdentityCode() string {
	return r._identityCode
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabasecurityjaqrpcloudrealnamecheckAPIRequest) GetImageUrls() string {
	return r._imageUrls
}
