package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页动态信息接口 API请求
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
type AlitripHotelDetailInfoGetRequest struct {
    model.Params
    // 详情页动态信息参数类
    paramHotelTopDetailsParam   *HotelTopDetailsParam
}

// 初始化AlitripHotelDetailInfoGetRequest对象
func NewAlitripHotelDetailInfoGetRequest() *AlitripHotelDetailInfoGetRequest{
    return &AlitripHotelDetailInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelDetailInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.detail.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelDetailInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelTopDetailsParam Setter
// 详情页动态信息参数类
func (r *AlitripHotelDetailInfoGetRequest) SetParamHotelTopDetailsParam(paramHotelTopDetailsParam *HotelTopDetailsParam) error {
    r.paramHotelTopDetailsParam = paramHotelTopDetailsParam
    r.Set("param_hotel_top_details_param", paramHotelTopDetailsParam)
    return nil
}

// ParamHotelTopDetailsParam Getter
func (r AlitripHotelDetailInfoGetRequest) GetParamHotelTopDetailsParam() *HotelTopDetailsParam {
    return r.paramHotelTopDetailsParam
}
