package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店城市列表展示 API请求
alitrip.merchant.galaxy.city.list

雅高酒店城市列表展示，并且首字母列出酒店城市
*/
type AlitripMerchantGalaxyCityListRequest struct {
    model.Params
    // 商家租户id
    _tenantKey   string
    // 0国内 1国外
    _domestic   int64
}

// 初始化AlitripMerchantGalaxyCityListRequest对象
func NewAlitripMerchantGalaxyCityListRequest() *AlitripMerchantGalaxyCityListRequest{
    return &AlitripMerchantGalaxyCityListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityListRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.city.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityListRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyCityListRequest) GetTenantKey() string {
    return r._tenantKey
}
// Domestic Setter
// 0国内 1国外
func (r *AlitripMerchantGalaxyCityListRequest) SetDomestic(_domestic int64) error {
    r._domestic = _domestic
    r.Set("domestic", _domestic)
    return nil
}

// Domestic Getter
func (r AlitripMerchantGalaxyCityListRequest) GetDomestic() int64 {
    return r._domestic
}
