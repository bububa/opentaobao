package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店详细信息搜索 APIRequest
alitrip.merchant.galaxy.hotel.detail.search

星河服务=获取雅高酒店详细信息
*/
type AlitripMerchantGalaxyHotelDetailSearchRequest struct {
    model.Params

    // 租户id
    tenantKey   string 

    // 酒店详情入参
    hotelDetailsParam   *HotelDetailsParam 

}

func NewAlitripMerchantGalaxyHotelDetailSearchRequest() *AlitripMerchantGalaxyHotelDetailSearchRequest{
    return &AlitripMerchantGalaxyHotelDetailSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.hotel.detail.search"
}

func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyHotelDetailSearchRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyHotelDetailSearchRequest) SetHotelDetailsParam(hotelDetailsParam *HotelDetailsParam) error {
    r.hotelDetailsParam = hotelDetailsParam
    r.Set("hotel_details_param", hotelDetailsParam)
    return nil
}

func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetHotelDetailsParam() *HotelDetailsParam {
    return r.hotelDetailsParam
}

