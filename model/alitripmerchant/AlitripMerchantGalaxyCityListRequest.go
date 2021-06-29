package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市列表展示 APIRequest
alitrip.merchant.galaxy.city.list

雅高酒店城市列表展示，并且首字母列出酒店城市
*/
type AlitripMerchantGalaxyCityListRequest struct {
    model.Params

    // 商家租户id
    tenantKey   string 

    // 0国内 1国外
    domestic   int64 

}

func NewAlitripMerchantGalaxyCityListRequest() *AlitripMerchantGalaxyCityListRequest{
    return &AlitripMerchantGalaxyCityListRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyCityListRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.city.list"
}

func (r AlitripMerchantGalaxyCityListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyCityListRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyCityListRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyCityListRequest) SetDomestic(domestic int64) error {
    r.domestic = domestic
    r.Set("domestic", domestic)
    return nil
}

func (r AlitripMerchantGalaxyCityListRequest) GetDomestic() int64 {
    return r.domestic
}

