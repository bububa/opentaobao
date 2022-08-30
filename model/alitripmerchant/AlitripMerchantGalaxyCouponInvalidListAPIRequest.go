package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCouponInvalidListAPIRequest 用户失效优惠券列表 API请求
// alitrip.merchant.galaxy.coupon.invalid.list
//
// 雅高小程序用户失效优惠券列表
type AlitripMerchantGalaxyCouponInvalidListAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyCouponInvalidListRequest 初始化AlitripMerchantGalaxyCouponInvalidListAPIRequest对象
func NewAlitripMerchantGalaxyCouponInvalidListRequest() *AlitripMerchantGalaxyCouponInvalidListAPIRequest {
	return &AlitripMerchantGalaxyCouponInvalidListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.coupon.invalid.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyCouponInvalidListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyCouponInvalidListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyCouponInvalidListAPIRequest) GetToken() string {
	return r._token
}
