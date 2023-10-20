package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberlogoutAPIRequest 星河-用户登出 API请求
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
type AlitripmerchantgalaxymemberlogoutAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
	// 用户登录token
	_token string
}

// NewAlitripmerchantgalaxymemberlogoutRequest 初始化AlitripmerchantgalaxymemberlogoutAPIRequest对象
func NewAlitripmerchantgalaxymemberlogoutRequest() *AlitripmerchantgalaxymemberlogoutAPIRequest {
	return &AlitripmerchantgalaxymemberlogoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberlogoutAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.logout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberlogoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberlogoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户信息
func (r *AlitripmerchantgalaxymemberlogoutAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberlogoutAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录token
func (r *AlitripmerchantgalaxymemberlogoutAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymemberlogoutAPIRequest) GetToken() string {
	return r._token
}
