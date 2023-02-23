package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthAccountinfoAPIRequest 授权账号信息 API请求
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
type AlibabaAlisportsPassportAuthAccountinfoAPIRequest struct {
	model.Params
	// 阿里体育业务KEY
	_alispAppKey string
	// 授权token，有效期2小时
	_ssoToken string
}

// NewAlibabaAlisportsPassportAuthAccountinfoRequest 初始化AlibabaAlisportsPassportAuthAccountinfoAPIRequest对象
func NewAlibabaAlisportsPassportAuthAccountinfoRequest() *AlibabaAlisportsPassportAuthAccountinfoAPIRequest {
	return &AlibabaAlisportsPassportAuthAccountinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAuthAccountinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.auth.accountinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAuthAccountinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAuthAccountinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 阿里体育业务KEY
func (r *AlibabaAlisportsPassportAuthAccountinfoAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAuthAccountinfoAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetSsoToken is SsoToken Setter
// 授权token，有效期2小时
func (r *AlibabaAlisportsPassportAuthAccountinfoAPIRequest) SetSsoToken(_ssoToken string) error {
	r._ssoToken = _ssoToken
	r.Set("sso_token", _ssoToken)
	return nil
}

// GetSsoToken SsoToken Getter
func (r AlibabaAlisportsPassportAuthAccountinfoAPIRequest) GetSsoToken() string {
	return r._ssoToken
}
