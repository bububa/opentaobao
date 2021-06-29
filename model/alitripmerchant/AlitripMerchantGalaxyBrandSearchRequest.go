package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-品牌搜索 API请求
alitrip.merchant.galaxy.brand.search

星河服务=获取雅高品牌信息
*/
type AlitripMerchantGalaxyBrandSearchRequest struct {
    model.Params
    // 租户信息
    _tenantKey   string
}

// 初始化AlitripMerchantGalaxyBrandSearchRequest对象
func NewAlitripMerchantGalaxyBrandSearchRequest() *AlitripMerchantGalaxyBrandSearchRequest{
    return &AlitripMerchantGalaxyBrandSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyBrandSearchRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.brand.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyBrandSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyBrandSearchRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyBrandSearchRequest) GetTenantKey() string {
    return r._tenantKey
}
