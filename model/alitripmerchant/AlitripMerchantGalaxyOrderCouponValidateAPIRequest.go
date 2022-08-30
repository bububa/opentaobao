package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderCouponValidateAPIRequest 携带券的试单接口 API请求
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
type AlitripMerchantGalaxyOrderCouponValidateAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 使用优惠券的入参
	_couponParam *CouponParam
}

// NewAlitripMerchantGalaxyOrderCouponValidateRequest 初始化AlitripMerchantGalaxyOrderCouponValidateAPIRequest对象
func NewAlitripMerchantGalaxyOrderCouponValidateRequest() *AlitripMerchantGalaxyOrderCouponValidateAPIRequest {
	return &AlitripMerchantGalaxyOrderCouponValidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderCouponValidateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.coupon.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderCouponValidateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderCouponValidateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderCouponValidateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyOrderCouponValidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderCouponValidateAPIRequest) GetToken() string {
	return r._token
}

// SetCouponParam is CouponParam Setter
// 使用优惠券的入参
func (r *AlitripMerchantGalaxyOrderCouponValidateAPIRequest) SetCouponParam(_couponParam *CouponParam) error {
	r._couponParam = _couponParam
	r.Set("coupon_param", _couponParam)
	return nil
}

// GetCouponParam CouponParam Getter
func (r AlitripMerchantGalaxyOrderCouponValidateAPIRequest) GetCouponParam() *CouponParam {
	return r._couponParam
}
