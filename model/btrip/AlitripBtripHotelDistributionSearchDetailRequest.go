package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-详情报价接口 API请求
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口
*/
type AlitripBtripHotelDistributionSearchDetailRequest struct {
    model.Params
    // 详情报价入参
    _paramHotelDetailRQ   *HotelDetailRq
}

// 初始化AlitripBtripHotelDistributionSearchDetailRequest对象
func NewAlitripBtripHotelDistributionSearchDetailRequest() *AlitripBtripHotelDistributionSearchDetailRequest{
    return &AlitripBtripHotelDistributionSearchDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchDetailRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.search.detail"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelDetailRQ Setter
// 详情报价入参
func (r *AlitripBtripHotelDistributionSearchDetailRequest) SetParamHotelDetailRQ(_paramHotelDetailRQ *HotelDetailRq) error {
    r._paramHotelDetailRQ = _paramHotelDetailRQ
    r.Set("param_hotel_detail_r_q", _paramHotelDetailRQ)
    return nil
}

// ParamHotelDetailRQ Getter
func (r AlitripBtripHotelDistributionSearchDetailRequest) GetParamHotelDetailRQ() *HotelDetailRq {
    return r._paramHotelDetailRQ
}
