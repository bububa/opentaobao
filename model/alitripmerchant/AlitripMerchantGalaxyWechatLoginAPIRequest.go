package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatloginAPIRequest 星河-用户使用微信登陆 API请求
// alitrip.merchant.galaxy.wechat.login
//
// 星河产品=用户微信小程序登陆
type AlitripmerchantgalaxywechatloginAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 微信小程序登陆请求参数
	_loginParam *LoginParam
}

// NewAlitripmerchantgalaxywechatloginRequest 初始化AlitripmerchantgalaxywechatloginAPIRequest对象
func NewAlitripmerchantgalaxywechatloginRequest() *AlitripmerchantgalaxywechatloginAPIRequest {
	return &AlitripmerchantgalaxywechatloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatloginAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxywechatloginAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatloginAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetLoginParam is LoginParam Setter
// 微信小程序登陆请求参数
func (r *AlitripmerchantgalaxywechatloginAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripmerchantgalaxywechatloginAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}
