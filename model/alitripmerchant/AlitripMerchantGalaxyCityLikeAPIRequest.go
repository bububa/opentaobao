package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCityLikeAPIRequest 星河-酒店城市模糊查询 API请求
// alitrip.merchant.galaxy.city.like
//
// 根据城市模糊查询，雅高酒店所在城市的城市信息
type AlitripMerchantGalaxyCityLikeAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 城市模糊
	_cityName string
	// 0国内1国外
	_domestic int64
}

// NewAlitripMerchantGalaxyCityLikeRequest 初始化AlitripMerchantGalaxyCityLikeAPIRequest对象
func NewAlitripMerchantGalaxyCityLikeRequest() *AlitripMerchantGalaxyCityLikeAPIRequest {
	return &AlitripMerchantGalaxyCityLikeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) Reset() {
	r._tenantKey = ""
	r._cityName = ""
	r._domestic = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.like"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCityName is CityName Setter
// 城市模糊
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// GetCityName CityName Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetCityName() string {
	return r._cityName
}

// SetDomestic is Domestic Setter
// 0国内1国外
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetDomestic() int64 {
	return r._domestic
}

var poolAlitripMerchantGalaxyCityLikeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyCityLikeRequest()
	},
}

// GetAlitripMerchantGalaxyCityLikeRequest 从 sync.Pool 获取 AlitripMerchantGalaxyCityLikeAPIRequest
func GetAlitripMerchantGalaxyCityLikeAPIRequest() *AlitripMerchantGalaxyCityLikeAPIRequest {
	return poolAlitripMerchantGalaxyCityLikeAPIRequest.Get().(*AlitripMerchantGalaxyCityLikeAPIRequest)
}

// ReleaseAlitripMerchantGalaxyCityLikeAPIRequest 将 AlitripMerchantGalaxyCityLikeAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyCityLikeAPIRequest(v *AlitripMerchantGalaxyCityLikeAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyCityLikeAPIRequest.Put(v)
}
