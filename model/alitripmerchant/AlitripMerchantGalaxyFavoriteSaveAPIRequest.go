package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyfavoritesaveAPIRequest 用户添加或移除收藏接口 API请求
// alitrip.merchant.galaxy.favorite.save
//
// 用户可以点击收藏酒店，再次点击移除收藏的酒店
type AlitripmerchantgalaxyfavoritesaveAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 酒店id
	_hotelId string
}

// NewAlitripmerchantgalaxyfavoritesaveRequest 初始化AlitripmerchantgalaxyfavoritesaveAPIRequest对象
func NewAlitripmerchantgalaxyfavoritesaveRequest() *AlitripmerchantgalaxyfavoritesaveAPIRequest {
	return &AlitripmerchantgalaxyfavoritesaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyfavoritesaveAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyfavoritesaveAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetToken() string {
	return r._token
}

// SetHotelId is HotelId Setter
// 酒店id
func (r *AlitripmerchantgalaxyfavoritesaveAPIRequest) SetHotelId(_hotelId string) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r AlitripmerchantgalaxyfavoritesaveAPIRequest) GetHotelId() string {
	return r._hotelId
}
