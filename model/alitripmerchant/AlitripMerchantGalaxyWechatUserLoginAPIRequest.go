package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatUserLoginAPIRequest 微信小程序用户登录 API请求
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
type AlitripMerchantGalaxyWechatUserLoginAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 微信向程序登录参数
	_loginParam *LoginParam
}

// NewAlitripMerchantGalaxyWechatUserLoginRequest 初始化AlitripMerchantGalaxyWechatUserLoginAPIRequest对象
func NewAlitripMerchantGalaxyWechatUserLoginRequest() *AlitripMerchantGalaxyWechatUserLoginAPIRequest {
	return &AlitripMerchantGalaxyWechatUserLoginAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatUserLoginAPIRequest) Reset() {
	r._tenantKey = ""
	r._loginParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatUserLoginAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.user.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatUserLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatUserLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyWechatUserLoginAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatUserLoginAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetLoginParam is LoginParam Setter
// 微信向程序登录参数
func (r *AlitripMerchantGalaxyWechatUserLoginAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripMerchantGalaxyWechatUserLoginAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}

var poolAlitripMerchantGalaxyWechatUserLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatUserLoginRequest()
	},
}

// GetAlitripMerchantGalaxyWechatUserLoginRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatUserLoginAPIRequest
func GetAlitripMerchantGalaxyWechatUserLoginAPIRequest() *AlitripMerchantGalaxyWechatUserLoginAPIRequest {
	return poolAlitripMerchantGalaxyWechatUserLoginAPIRequest.Get().(*AlitripMerchantGalaxyWechatUserLoginAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatUserLoginAPIRequest 将 AlitripMerchantGalaxyWechatUserLoginAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatUserLoginAPIRequest(v *AlitripMerchantGalaxyWechatUserLoginAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatUserLoginAPIRequest.Put(v)
}
