package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyfavoritelistAPIRequest 用户收藏列表查询 API请求
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
type AlitripmerchantgalaxyfavoritelistAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
}

// NewAlitripmerchantgalaxyfavoritelistRequest 初始化AlitripmerchantgalaxyfavoritelistAPIRequest对象
func NewAlitripmerchantgalaxyfavoritelistRequest() *AlitripmerchantgalaxyfavoritelistAPIRequest {
	return &AlitripmerchantgalaxyfavoritelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyfavoritelistAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyfavoritelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyfavoritelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyfavoritelistAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyfavoritelistAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyfavoritelistAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyfavoritelistAPIRequest) GetToken() string {
	return r._token
}
