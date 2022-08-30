package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityCouponListAPIRequest 用户领券中心列表 API请求
// alitrip.merchant.galaxy.activity.coupon.list
//
// 雅高小程序用户领券中心列表
type AlitripMerchantGalaxyActivityCouponListAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户Token
	_token string
}

// NewAlitripMerchantGalaxyActivityCouponListRequest 初始化AlitripMerchantGalaxyActivityCouponListAPIRequest对象
func NewAlitripMerchantGalaxyActivityCouponListRequest() *AlitripMerchantGalaxyActivityCouponListAPIRequest {
	return &AlitripMerchantGalaxyActivityCouponListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.coupon.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripMerchantGalaxyActivityCouponListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户Token
func (r *AlitripMerchantGalaxyActivityCouponListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityCouponListAPIRequest) GetToken() string {
	return r._token
}
