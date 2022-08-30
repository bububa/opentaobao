package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest 按活动Id领取优惠券 API请求
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
type AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户token
	_token string
	// 活动Id集合(逗号分隔)
	_activityIds string
}

// NewAlitripMerchantGalaxyReceiveCouponByActivityRequest 初始化AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest对象
func NewAlitripMerchantGalaxyReceiveCouponByActivityRequest() *AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest {
	return &AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.receive.coupon.by.activity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) GetToken() string {
	return r._token
}

// SetActivityIds is ActivityIds Setter
// 活动Id集合(逗号分隔)
func (r *AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) SetActivityIds(_activityIds string) error {
	r._activityIds = _activityIds
	r.Set("activity_ids", _activityIds)
	return nil
}

// GetActivityIds ActivityIds Getter
func (r AlitripMerchantGalaxyReceiveCouponByActivityAPIRequest) GetActivityIds() string {
	return r._activityIds
}
