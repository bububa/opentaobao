package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest sso_token验证 API请求
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest struct {
	model.Params
	// sso_token
	_ssoToken string
	// 应用APPKEY
	_alispAppKey string
	// 当前时间戳[精确到秒，10位]
	_alispTime string
	// 签名
	_alispSign string
}

// NewAlibabaAlisportsPassportAccountSsotokenvalidateRequest 初始化AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest对象
func NewAlibabaAlisportsPassportAccountSsotokenvalidateRequest() *AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest {
	return &AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.ssotokenvalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSsoToken is SsoToken Setter
// sso_token
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) SetSsoToken(_ssoToken string) error {
	r._ssoToken = _ssoToken
	r.Set("sso_token", _ssoToken)
	return nil
}

// GetSsoToken SsoToken Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetSsoToken() string {
	return r._ssoToken
}

// SetAlispAppKey is AlispAppKey Setter
// 应用APPKEY
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 当前时间戳[精确到秒，10位]
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateAPIRequest) GetAlispSign() string {
	return r._alispSign
}
