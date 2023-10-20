package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyordercouponvalidateAPIRequest 携带券的试单接口 API请求
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
type AlitripmerchantgalaxyordercouponvalidateAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 使用优惠券的入参
	_couponParam *CouponParam
}

// NewAlitripmerchantgalaxyordercouponvalidateRequest 初始化AlitripmerchantgalaxyordercouponvalidateAPIRequest对象
func NewAlitripmerchantgalaxyordercouponvalidateRequest() *AlitripmerchantgalaxyordercouponvalidateAPIRequest {
	return &AlitripmerchantgalaxyordercouponvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.coupon.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyordercouponvalidateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyordercouponvalidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetToken() string {
	return r._token
}

// SetCouponParam is CouponParam Setter
// 使用优惠券的入参
func (r *AlitripmerchantgalaxyordercouponvalidateAPIRequest) SetCouponParam(_couponParam *CouponParam) error {
	r._couponParam = _couponParam
	r.Set("coupon_param", _couponParam)
	return nil
}

// GetCouponParam CouponParam Getter
func (r AlitripmerchantgalaxyordercouponvalidateAPIRequest) GetCouponParam() *CouponParam {
	return r._couponParam
}
