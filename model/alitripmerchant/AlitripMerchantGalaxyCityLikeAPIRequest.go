package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyCityLikeAPIRequest
星河-酒店城市模糊查询 API请求
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.like"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// Get TenantKey Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// Set is CityName Setter
// 城市模糊
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// Get CityName Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetCityName() string {
	return r._cityName
}

// Set is Domestic Setter
// 0国内1国外
func (r *AlitripMerchantGalaxyCityLikeAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// Get Domestic Getter
func (r AlitripMerchantGalaxyCityLikeAPIRequest) GetDomestic() int64 {
	return r._domestic
}
