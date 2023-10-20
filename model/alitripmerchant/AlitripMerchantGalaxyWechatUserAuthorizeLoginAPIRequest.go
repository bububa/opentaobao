package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) Reset() {
	r._tenantKey = ""
	r._loginParam = nil
	r.Params.ToZero()
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

var poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatUserAuthorizeLoginRequest()
	},
}

// GetAlitripMerchantGalaxyWechatUserAuthorizeLoginRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest
func GetAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest() *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest {
	return poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest.Get().(*AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest 将 AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest(v *AlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserAuthorizeLoginAPIRequest.Put(v)
}
