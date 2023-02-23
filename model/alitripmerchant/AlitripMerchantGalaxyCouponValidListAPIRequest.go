package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCouponValidListAPIRequest 用户有效优惠券列表 API请求
// alitrip.merchant.galaxy.coupon.valid.list
//
// 雅高小程序用户有效优惠券列表
type AlitripMerchantGalaxyCouponValidListAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户Token
	_token string
}

// NewAlitripMerchantGalaxyCouponValidListRequest 初始化AlitripMerchantGalaxyCouponValidListAPIRequest对象
func NewAlitripMerchantGalaxyCouponValidListRequest() *AlitripMerchantGalaxyCouponValidListAPIRequest {
	return &AlitripMerchantGalaxyCouponValidListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCouponValidListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.coupon.valid.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCouponValidListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCouponValidListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripMerchantGalaxyCouponValidListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCouponValidListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户Token
func (r *AlitripMerchantGalaxyCouponValidListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyCouponValidListAPIRequest) GetToken() string {
	return r._token
}
