package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店列表页搜索 APIRequest
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

func NewAlitripMerchantGalaxyHotelListSearchRequest() *AlitripMerchantGalaxyHotelListSearchRequest{
    return &AlitripMerchantGalaxyHotelListSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyHotelListSearchRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.hotel.list.search"
}

func (r AlitripMerchantGalaxyHotelListSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyHotelListSearchRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyHotelListSearchRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyHotelListSearchRequest) SetListSearchParam(listSearchParam *ListSearchParam) error {
    r.listSearchParam = listSearchParam
    r.Set("list_search_param", listSearchParam)
    return nil
}

func (r AlitripMerchantGalaxyHotelListSearchRequest) GetListSearchParam() *ListSearchParam {
    return r.listSearchParam
}

