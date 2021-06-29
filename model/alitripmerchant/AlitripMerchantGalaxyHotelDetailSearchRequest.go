package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店详细信息搜索 API请求
alitrip.merchant.galaxy.hotel.detail.search

星河服务=获取雅高酒店详细信息
*/
type AlitripMerchantGalaxyHotelDetailSearchRequest struct {
    model.Params
    // 租户id
    _tenantKey   string
    // 酒店详情入参
    _hotelDetailsParam   *HotelDetailsParam
}

// 初始化AlitripMerchantGalaxyHotelDetailSearchRequest对象
func NewAlitripMerchantGalaxyHotelDetailSearchRequest() *AlitripMerchantGalaxyHotelDetailSearchRequest{
    return &AlitripMerchantGalaxyHotelDetailSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.hotel.detail.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyHotelDetailSearchRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetTenantKey() string {
    return r._tenantKey
}
// HotelDetailsParam Setter
// 酒店详情入参
func (r *AlitripMerchantGalaxyHotelDetailSearchRequest) SetHotelDetailsParam(_hotelDetailsParam *HotelDetailsParam) error {
    r._hotelDetailsParam = _hotelDetailsParam
    r.Set("hotel_details_param", _hotelDetailsParam)
    return nil
}

// HotelDetailsParam Getter
func (r AlitripMerchantGalaxyHotelDetailSearchRequest) GetHotelDetailsParam() *HotelDetailsParam {
    return r._hotelDetailsParam
}
