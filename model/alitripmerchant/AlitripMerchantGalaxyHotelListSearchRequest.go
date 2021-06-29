package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店列表页搜索 API请求
alitrip.merchant.galaxy.hotel.list.search

星河产品=酒店列表页搜索
*/
type AlitripMerchantGalaxyHotelListSearchRequest struct {
    model.Params
    // 商家租户id
    tenantKey   string
    // 请求参数
    listSearchParam   *ListSearchParam
}

// 初始化AlitripMerchantGalaxyHotelListSearchRequest对象
func NewAlitripMerchantGalaxyHotelListSearchRequest() *AlitripMerchantGalaxyHotelListSearchRequest{
    return &AlitripMerchantGalaxyHotelListSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyHotelListSearchRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.hotel.list.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyHotelListSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyHotelListSearchRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyHotelListSearchRequest) GetTenantKey() string {
    return r.tenantKey
}
// ListSearchParam Setter
// 请求参数
func (r *AlitripMerchantGalaxyHotelListSearchRequest) SetListSearchParam(listSearchParam *ListSearchParam) error {
    r.listSearchParam = listSearchParam
    r.Set("list_search_param", listSearchParam)
    return nil
}

// ListSearchParam Getter
func (r AlitripMerchantGalaxyHotelListSearchRequest) GetListSearchParam() *ListSearchParam {
    return r.listSearchParam
}
