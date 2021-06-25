package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页静态信息 APIRequest
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
type AlitripHotelDetailStaticinfoGetRequest struct {
    model.Params

    // 酒店详情页静态信息参数
    paramHotelTopStaticDetailsParam   *HotelTopStaticDetailsParam 

}

func NewAlitripHotelDetailStaticinfoGetRequest() *AlitripHotelDetailStaticinfoGetRequest{
    return &AlitripHotelDetailStaticinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelDetailStaticinfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.detail.staticinfo.get"
}

func (r AlitripHotelDetailStaticinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelDetailStaticinfoGetRequest) SetParamHotelTopStaticDetailsParam(paramHotelTopStaticDetailsParam *HotelTopStaticDetailsParam) error {
    r.paramHotelTopStaticDetailsParam = paramHotelTopStaticDetailsParam
    r.Set("param_hotel_top_static_details_param", paramHotelTopStaticDetailsParam)
    return nil
}

func (r AlitripHotelDetailStaticinfoGetRequest) GetParamHotelTopStaticDetailsParam() *HotelTopStaticDetailsParam {
    return r.paramHotelTopStaticDetailsParam
}

