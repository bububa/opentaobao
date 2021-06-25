package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页动态信息接口 APIRequest
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
type AlitripHotelDetailInfoGetRequest struct {
    model.Params

    // 详情页动态信息参数类
    paramHotelTopDetailsParam   *HotelTopDetailsParam 

}

func NewAlitripHotelDetailInfoGetRequest() *AlitripHotelDetailInfoGetRequest{
    return &AlitripHotelDetailInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelDetailInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.detail.info.get"
}

func (r AlitripHotelDetailInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelDetailInfoGetRequest) SetParamHotelTopDetailsParam(paramHotelTopDetailsParam *HotelTopDetailsParam) error {
    r.paramHotelTopDetailsParam = paramHotelTopDetailsParam
    r.Set("param_hotel_top_details_param", paramHotelTopDetailsParam)
    return nil
}

func (r AlitripHotelDetailInfoGetRequest) GetParamHotelTopDetailsParam() *HotelTopDetailsParam {
    return r.paramHotelTopDetailsParam
}

