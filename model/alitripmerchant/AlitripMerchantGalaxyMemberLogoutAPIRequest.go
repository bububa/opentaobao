package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyMemberLogoutAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.logout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyMemberLogoutAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录token
func (r *AlitripMerchantGalaxyMemberLogoutAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyMemberLogoutAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyMemberLogoutAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyMemberLogoutRequest()
	},
}

// GetAlitripMerchantGalaxyMemberLogoutRequest 从 sync.Pool 获取 AlitripMerchantGalaxyMemberLogoutAPIRequest
func GetAlitripMerchantGalaxyMemberLogoutAPIRequest() *AlitripMerchantGalaxyMemberLogoutAPIRequest {
	return poolAlitripMerchantGalaxyMemberLogoutAPIRequest.Get().(*AlitripMerchantGalaxyMemberLogoutAPIRequest)
}

// ReleaseAlitripMerchantGalaxyMemberLogoutAPIRequest 将 AlitripMerchantGalaxyMemberLogoutAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyMemberLogoutAPIRequest(v *AlitripMerchantGalaxyMemberLogoutAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyMemberLogoutAPIRequest.Put(v)
}
