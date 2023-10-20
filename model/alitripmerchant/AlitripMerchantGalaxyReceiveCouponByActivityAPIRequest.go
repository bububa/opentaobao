package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest 按活动Id领取优惠券 API请求
// alitrip.merchant.galaxy.receive.coupon.by.activity
//
// 雅高小程序按活动Id领取优惠券
type AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户token
	_token string
	// 活动Id集合(逗号分隔)
	_activityIds string
}

// NewAlitripmerchantgalaxyreceivecouponbyactivityRequest 初始化AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest对象
func NewAlitripmerchantgalaxyreceivecouponbyactivityRequest() *AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest {
	return &AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.receive.coupon.by.activity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetToken() string {
	return r._token
}

// SetActivityIds is ActivityIds Setter
// 活动Id集合(逗号分隔)
func (r *AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) SetActivityIds(_activityIds string) error {
	r._activityIds = _activityIds
	r.Set("activity_ids", _activityIds)
	return nil
}

// GetActivityIds ActivityIds Getter
func (r AlitripmerchantgalaxyreceivecouponbyactivityAPIRequest) GetActivityIds() string {
	return r._activityIds
}
