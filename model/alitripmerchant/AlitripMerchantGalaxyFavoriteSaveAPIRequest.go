package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteSaveAPIRequest 用户添加或移除收藏接口 API请求
// alitrip.merchant.galaxy.favorite.save
//
// 用户可以点击收藏酒店，再次点击移除收藏的酒店
type AlitripMerchantGalaxyFavoriteSaveAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 酒店id
	_hotelId string
}

// NewAlitripMerchantGalaxyFavoriteSaveRequest 初始化AlitripMerchantGalaxyFavoriteSaveAPIRequest对象
func NewAlitripMerchantGalaxyFavoriteSaveRequest() *AlitripMerchantGalaxyFavoriteSaveAPIRequest {
	return &AlitripMerchantGalaxyFavoriteSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyFavoriteSaveAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyFavoriteSaveAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetToken() string {
	return r._token
}

// SetHotelId is HotelId Setter
// 酒店id
func (r *AlitripMerchantGalaxyFavoriteSaveAPIRequest) SetHotelId(_hotelId string) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r AlitripMerchantGalaxyFavoriteSaveAPIRequest) GetHotelId() string {
	return r._hotelId
}
