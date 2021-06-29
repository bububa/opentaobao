package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-酒店最低价 APIRequest
alitrip.btrip.hotel.distribution.search.low.price

商旅酒店api分销-酒店最低价
*/
type AlitripBtripHotelDistributionSearchLowPriceRequest struct {
    model.Params

    // 列表最低价入参
    paramHotelSearchListRQ   *HotelSearchListRq 

}

func NewAlitripBtripHotelDistributionSearchLowPriceRequest() *AlitripBtripHotelDistributionSearchLowPriceRequest{
    return &AlitripBtripHotelDistributionSearchLowPriceRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.search.low.price"
}

func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionSearchLowPriceRequest) SetParamHotelSearchListRQ(paramHotelSearchListRQ *HotelSearchListRq) error {
    r.paramHotelSearchListRQ = paramHotelSearchListRQ
    r.Set("param_hotel_search_list_r_q", paramHotelSearchListRQ)
    return nil
}

func (r AlitripBtripHotelDistributionSearchLowPriceRequest) GetParamHotelSearchListRQ() *HotelSearchListRq {
    return r.paramHotelSearchListRQ
}

