package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycouponinvalidlistAPIRequest 用户失效优惠券列表 API请求
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
type AlitripmerchantgalaxycouponinvalidlistAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripmerchantgalaxycouponinvalidlistRequest 初始化AlitripmerchantgalaxycouponinvalidlistAPIRequest对象
func NewAlitripmerchantgalaxycouponinvalidlistRequest() *AlitripmerchantgalaxycouponinvalidlistAPIRequest {
	return &AlitripmerchantgalaxycouponinvalidlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycouponinvalidlistAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.coupon.invalid.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycouponinvalidlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycouponinvalidlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxycouponinvalidlistAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycouponinvalidlistAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxycouponinvalidlistAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxycouponinvalidlistAPIRequest) GetToken() string {
	return r._token
}
