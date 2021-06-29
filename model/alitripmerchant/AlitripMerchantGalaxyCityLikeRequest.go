package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市模糊查询 API请求
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息
*/
type AlitripMerchantGalaxyCityLikeRequest struct {
    model.Params
    // 商家租户id
    _tenantKey   string
    // 城市模糊
    _cityName   string
    // 0国内1国外
    _domestic   int64
}

// 初始化AlitripMerchantGalaxyCityLikeRequest对象
func NewAlitripMerchantGalaxyCityLikeRequest() *AlitripMerchantGalaxyCityLikeRequest{
    return &AlitripMerchantGalaxyCityLikeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityLikeRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.city.like"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityLikeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityLikeRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyCityLikeRequest) GetTenantKey() string {
    return r._tenantKey
}
// CityName Setter
// 城市模糊
func (r *AlitripMerchantGalaxyCityLikeRequest) SetCityName(_cityName string) error {
    r._cityName = _cityName
    r.Set("city_name", _cityName)
    return nil
}

// CityName Getter
func (r AlitripMerchantGalaxyCityLikeRequest) GetCityName() string {
    return r._cityName
}
// Domestic Setter
// 0国内1国外
func (r *AlitripMerchantGalaxyCityLikeRequest) SetDomestic(_domestic int64) error {
    r._domestic = _domestic
    r.Set("domestic", _domestic)
    return nil
}

// Domestic Getter
func (r AlitripMerchantGalaxyCityLikeRequest) GetDomestic() int64 {
    return r._domestic
}
