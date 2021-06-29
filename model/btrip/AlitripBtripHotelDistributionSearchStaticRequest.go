package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-酒店静态信息接口 API请求
alitrip.btrip.hotel.distribution.search.static

商旅酒店api分销-酒店静态信息接口
*/
type AlitripBtripHotelDistributionSearchStaticRequest struct {
    model.Params
    // 基础信息入参
    paramHotelInfoRQ   *HotelInfoRq
}

// 初始化AlitripBtripHotelDistributionSearchStaticRequest对象
func NewAlitripBtripHotelDistributionSearchStaticRequest() *AlitripBtripHotelDistributionSearchStaticRequest{
    return &AlitripBtripHotelDistributionSearchStaticRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchStaticRequest) GetApiMethodName() string {
    return "alitrip.btrip.hotel.distribution.search.static"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchStaticRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelInfoRQ Setter
// 基础信息入参
func (r *AlitripBtripHotelDistributionSearchStaticRequest) SetParamHotelInfoRQ(paramHotelInfoRQ *HotelInfoRq) error {
    r.paramHotelInfoRQ = paramHotelInfoRQ
    r.Set("param_hotel_info_r_q", paramHotelInfoRQ)
    return nil
}

// ParamHotelInfoRQ Getter
func (r AlitripBtripHotelDistributionSearchStaticRequest) GetParamHotelInfoRQ() *HotelInfoRq {
    return r.paramHotelInfoRQ
}
