package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyfavoritequeryAPIRequest 单酒店收藏状态查询 API请求
// alitrip.merchant.galaxy.favorite.query
//
// 返回用户对单个酒店的收藏状态
type AlitripmerchantgalaxyfavoritequeryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
	// 酒店id
	_hotelId string
}

// NewAlitripmerchantgalaxyfavoritequeryRequest 初始化AlitripmerchantgalaxyfavoritequeryAPIRequest对象
func NewAlitripmerchantgalaxyfavoritequeryRequest() *AlitripmerchantgalaxyfavoritequeryAPIRequest {
	return &AlitripmerchantgalaxyfavoritequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxyfavoritequeryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripmerchantgalaxyfavoritequeryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetToken() string {
	return r._token
}

// SetHotelId is HotelId Setter
// 酒店id
func (r *AlitripmerchantgalaxyfavoritequeryAPIRequest) SetHotelId(_hotelId string) error {
	r._hotelId = _hotelId
	r.Set("hotel_id", _hotelId)
	return nil
}

// GetHotelId HotelId Getter
func (r AlitripmerchantgalaxyfavoritequeryAPIRequest) GetHotelId() string {
	return r._hotelId
}
