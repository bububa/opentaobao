package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatLoginAPIRequest 星河-用户使用微信登陆 API请求
// alitrip.merchant.galaxy.wechat.login
//
// 星河产品=用户微信小程序登陆
type AlitripMerchantGalaxyWechatLoginAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 微信小程序登陆请求参数
	_loginParam *LoginParam
}

// NewAlitripMerchantGalaxyWechatLoginRequest 初始化AlitripMerchantGalaxyWechatLoginAPIRequest对象
func NewAlitripMerchantGalaxyWechatLoginRequest() *AlitripMerchantGalaxyWechatLoginAPIRequest {
	return &AlitripMerchantGalaxyWechatLoginAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatLoginAPIRequest) Reset() {
	r._tenantKey = ""
	r._loginParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyWechatLoginAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetLoginParam is LoginParam Setter
// 微信小程序登陆请求参数
func (r *AlitripMerchantGalaxyWechatLoginAPIRequest) SetLoginParam(_loginParam *LoginParam) error {
	r._loginParam = _loginParam
	r.Set("login_param", _loginParam)
	return nil
}

// GetLoginParam LoginParam Getter
func (r AlitripMerchantGalaxyWechatLoginAPIRequest) GetLoginParam() *LoginParam {
	return r._loginParam
}

var poolAlitripMerchantGalaxyWechatLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatLoginRequest()
	},
}

// GetAlitripMerchantGalaxyWechatLoginRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatLoginAPIRequest
func GetAlitripMerchantGalaxyWechatLoginAPIRequest() *AlitripMerchantGalaxyWechatLoginAPIRequest {
	return poolAlitripMerchantGalaxyWechatLoginAPIRequest.Get().(*AlitripMerchantGalaxyWechatLoginAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatLoginAPIRequest 将 AlitripMerchantGalaxyWechatLoginAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatLoginAPIRequest(v *AlitripMerchantGalaxyWechatLoginAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatLoginAPIRequest.Put(v)
}
