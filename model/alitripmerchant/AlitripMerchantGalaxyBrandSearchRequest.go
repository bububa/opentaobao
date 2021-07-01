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
type AlitripMerchantGalaxyBrandSearchAPIRequest struct {
    model.Params
    // 租户信息
    _tenantKey   string
}

// 初始化AlitripMerchantGalaxyBrandSearchAPIRequest对象
func NewAlitripMerchantGalaxyBrandSearchRequest() *AlitripMerchantGalaxyBrandSearchAPIRequest{
    return &AlitripMerchantGalaxyBrandSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.brand.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户信息
func (r *AlitripMerchantGalaxyBrandSearchAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyBrandSearchAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
