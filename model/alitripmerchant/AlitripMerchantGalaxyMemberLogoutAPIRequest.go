package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberLogoutAPIRequest 星河-用户登出 API请求
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
type AlitripMerchantGalaxyMemberLogoutAPIRequest struct {
	model.Params
	// 租户信息
	_tenantKey string
	// 用户登录token
	_token string
}

// NewAlitripMerchantGalaxyMemberLogoutRequest 初始化AlitripMerchantGalaxyMemberLogoutAPIRequest对象
func NewAlitripMerchantGalaxyMemberLogoutRequest() *AlitripMerchantGalaxyMemberLogoutAPIRequest {
	return &AlitripMerchantGalaxyMemberLogoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.logout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyMemberLogoutAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyMemberLogoutAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetToken() string {
	return r._token
}
