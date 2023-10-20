package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivityfatigueAPIRequest 营销抽奖-弹窗疲劳度控制 API请求
// alitrip.merchant.galaxy.activity.fatigue
//
// 星河产品-营销抽奖首页弹窗疲劳度控制服务
type AlitripmerchantgalaxyactivityfatigueAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户登录标识
	_token string
	// 活动ID
	_offerId int64
}

// NewAlitripmerchantgalaxyactivityfatigueRequest 初始化AlitripmerchantgalaxyactivityfatigueAPIRequest对象
func NewAlitripmerchantgalaxyactivityfatigueRequest() *AlitripmerchantgalaxyactivityfatigueAPIRequest {
	return &AlitripmerchantgalaxyactivityfatigueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.fatigue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxyactivityfatigueAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyactivityfatigueAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动ID
func (r *AlitripmerchantgalaxyactivityfatigueAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripmerchantgalaxyactivityfatigueAPIRequest) GetOfferId() int64 {
	return r._offerId
}
