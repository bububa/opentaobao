package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest DFC-ID用户手机号授权登录 API请求
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
type AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 微信向程序登录参数
	_loginParam *LoginParam
}

// NewAlitripMerchantGalaxyWechatUserAuthorizeLoginRequest 初始化AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest对象
func NewAlitripMerchantGalaxyWechatUserAuthorizeLoginRequest() *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest {
	return &AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.user.authorize.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetLoginParam is LoginParam Setter
// 微信向程序登录参数
func (r *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}
