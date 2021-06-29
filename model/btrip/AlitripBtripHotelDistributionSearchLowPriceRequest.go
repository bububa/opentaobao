package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-酒店最低价 API请求
alitrip.btrip.hotel.distribution.search.low.price

商旅酒店api分销-酒店最低价
*/
type AlitripBtripHotelDistributionSearchLowPriceRequest struct {
    model.Params
    // 列表最低价入参
    paramHotelSearchListRQ   *HotelSearchListRq
}

// 初始化AlitripBtripHotelDistributionSearchLowPriceRequest对象
func NewAlitripBtripHotelDistributionSearchLowPriceRequest() *AlitripBtripHotelDistributionSearchLowPriceRequest{
    return &AlitripBtripHotelDistributionSearchLowPriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.search.low.price"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelSearchListRQ Setter
// 列表最低价入参
func (r *AlitripBtripHotelDistributionSearchLowPriceRequest) SetParamHotelSearchListRQ(paramHotelSearchListRQ *HotelSearchListRq) error {
    r.paramHotelSearchListRQ = paramHotelSearchListRQ
    r.Set("param_hotel_search_list_r_q", paramHotelSearchListRQ)
    return nil
}

// ParamHotelSearchListRQ Getter
func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetParamHotelSearchListRQ() *HotelSearchListRq {
    return r.paramHotelSearchListRQ
}
