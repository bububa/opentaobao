package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市模糊查询 APIRequest
alitrip.merchant.galaxy.city.like

根据城市模糊查询，雅高酒店所在城市的城市信息
*/
type AlitripMerchantGalaxyCityLikeRequest struct {
    model.Params

    // 商家租户id
    tenantKey   string 

    // 城市模糊
    cityName   string 

    // 0国内1国外
    domestic   int64 

}

func NewAlitripMerchantGalaxyCityLikeRequest() *AlitripMerchantGalaxyCityLikeRequest{
    return &AlitripMerchantGalaxyCityLikeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyCityLikeRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.city.like"
}

func (r AlitripMerchantGalaxyCityLikeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyCityLikeRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyCityLikeRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyCityLikeRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

func (r AlitripMerchantGalaxyCityLikeRequest) GetCityName() string {
    return r.cityName
}

func (r *AlitripMerchantGalaxyCityLikeRequest) SetDomestic(domestic int64) error {
    r.domestic = domestic
    r.Set("domestic", domestic)
    return nil
}

func (r AlitripMerchantGalaxyCityLikeRequest) GetDomestic() int64 {
    return r.domestic
}

