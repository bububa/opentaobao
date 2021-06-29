package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-详情报价接口 APIRequest
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口
*/
type AlitripBtripHotelDistributionSearchDetailRequest struct {
    model.Params

    // 详情报价入参
    paramHotelDetailRQ   *HotelDetailRq 

}

func NewAlitripBtripHotelDistributionSearchDetailRequest() *AlitripBtripHotelDistributionSearchDetailRequest{
    return &AlitripBtripHotelDistributionSearchDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripHotelDistributionSearchDetailRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.search.detail"
}

func (r AlitripBtripHotelDistributionSearchDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripHotelDistributionSearchDetailRequest) SetParamHotelDetailRQ(paramHotelDetailRQ *HotelDetailRq) error {
    r.paramHotelDetailRQ = paramHotelDetailRQ
    r.Set("param_hotel_detail_r_q", paramHotelDetailRQ)
    return nil
}

func (r AlitripBtripHotelDistributionSearchDetailRequest) GetParamHotelDetailRQ() *HotelDetailRq {
    return r.paramHotelDetailRQ
}

