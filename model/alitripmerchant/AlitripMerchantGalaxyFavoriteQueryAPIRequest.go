package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyFavoriteQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._hotelId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyFavoriteQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripMerchantGalaxyFavoriteQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyFavoriteQueryRequest()
	},
}

// GetAlitripMerchantGalaxyFavoriteQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyFavoriteQueryAPIRequest
func GetAlitripMerchantGalaxyFavoriteQueryAPIRequest() *AlitripMerchantGalaxyFavoriteQueryAPIRequest {
	return poolAlitripMerchantGalaxyFavoriteQueryAPIRequest.Get().(*AlitripMerchantGalaxyFavoriteQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyFavoriteQueryAPIRequest 将 AlitripMerchantGalaxyFavoriteQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyFavoriteQueryAPIRequest(v *AlitripMerchantGalaxyFavoriteQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyFavoriteQueryAPIRequest.Put(v)
}
