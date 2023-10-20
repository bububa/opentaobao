package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitypopupcontrolAPIRequest 营销弹屏疲劳度控制 API请求
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
type AlitripmerchantgalaxyactivitypopupcontrolAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 活动类型
	_type string
	// 活动id
	_offerId int64
}

// NewAlitripmerchantgalaxyactivitypopupcontrolRequest 初始化AlitripmerchantgalaxyactivitypopupcontrolAPIRequest对象
func NewAlitripmerchantgalaxyactivitypopupcontrolRequest() *AlitripmerchantgalaxyactivitypopupcontrolAPIRequest {
	return &AlitripmerchantgalaxyactivitypopupcontrolAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.popup.control"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetToken() string {
	return r._token
}

// SetType is Type Setter
// 活动类型
func (r *AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetType() string {
	return r._type
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripmerchantgalaxyactivitypopupcontrolAPIRequest) GetOfferId() int64 {
	return r._offerId
}
