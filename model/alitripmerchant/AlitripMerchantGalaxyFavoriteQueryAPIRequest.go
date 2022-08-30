package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteQueryAPIRequest 单酒店收藏状态查询 API请求
// alitrip.merchant.galaxy.favorite.query
//
// 返回用户对单个酒店的收藏状态
type AlitripMerchantGalaxyFavoriteQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 酒店id
	_hotelId string
}

// NewAlitripMerchantGalaxyFavoriteQueryRequest 初始化AlitripMerchantGalaxyFavoriteQueryAPIRequest对象
func NewAlitripMerchantGalaxyFavoriteQueryRequest() *AlitripMerchantGalaxyFavoriteQueryAPIRequest {
	return &AlitripMerchantGalaxyFavoriteQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyFavoriteQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyFavoriteQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetToken() string {
	return r._token
}

// SetHotelId is HotelId Setter
// 酒店id
func (r *AlitripMerchantGalaxyFavoriteQueryAPIRequest) SetHotelId(_hotelId string) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetHotelId() string {
	return r._hotelId
}
