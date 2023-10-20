package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitycouponlistAPIRequest 用户领券中心列表 API请求
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
type AlitripmerchantgalaxyactivitycouponlistAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户Token
	_token string
}

// NewAlitripmerchantgalaxyactivitycouponlistRequest 初始化AlitripmerchantgalaxyactivitycouponlistAPIRequest对象
func NewAlitripmerchantgalaxyactivitycouponlistRequest() *AlitripmerchantgalaxyactivitycouponlistAPIRequest {
	return &AlitripmerchantgalaxyactivitycouponlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitycouponlistAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.coupon.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitycouponlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitycouponlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripmerchantgalaxyactivitycouponlistAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitycouponlistAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户Token
func (r *AlitripmerchantgalaxyactivitycouponlistAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitycouponlistAPIRequest) GetToken() string {
	return r._token
}
