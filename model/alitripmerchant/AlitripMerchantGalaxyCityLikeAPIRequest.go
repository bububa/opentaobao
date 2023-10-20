package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycitylikeAPIRequest 星河-酒店城市模糊查询 API请求
// alitrip.merchant.galaxy.city.like
//
// 根据城市模糊查询，雅高酒店所在城市的城市信息
type AlitripmerchantgalaxycitylikeAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 城市模糊
	_cityName string
	// 0国内1国外
	_domestic int64
}

// NewAlitripmerchantgalaxycitylikeRequest 初始化AlitripmerchantgalaxycitylikeAPIRequest对象
func NewAlitripmerchantgalaxycitylikeRequest() *AlitripmerchantgalaxycitylikeAPIRequest {
	return &AlitripmerchantgalaxycitylikeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.like"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxycitylikeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetCityName is CityName Setter
// 城市模糊
func (r *AlitripmerchantgalaxycitylikeAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// GetCityName CityName Getter
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetCityName() string {
	return r._cityName
}

// SetDomestic is Domestic Setter
// 0国内1国外
func (r *AlitripmerchantgalaxycitylikeAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r AlitripmerchantgalaxycitylikeAPIRequest) GetDomestic() int64 {
	return r._domestic
}
