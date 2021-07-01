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
type AlitripHotelDetailInfoGetAPIRequest struct {
    model.Params
    // 详情页动态信息参数类
    _paramHotelTopDetailsParam   *HotelTopDetailsParam
}

// 初始化AlitripHotelDetailInfoGetAPIRequest对象
func NewAlitripHotelDetailInfoGetRequest() *AlitripHotelDetailInfoGetAPIRequest{
    return &AlitripHotelDetailInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelDetailInfoGetAPIRequest) GetApiMethodName() string {
    return "alitrip.hotel.detail.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelDetailInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelTopDetailsParam Setter
// 详情页动态信息参数类
func (r *AlitripHotelDetailInfoGetAPIRequest) SetParamHotelTopDetailsParam(_paramHotelTopDetailsParam *HotelTopDetailsParam) error {
    r._paramHotelTopDetailsParam = _paramHotelTopDetailsParam
    r.Set("param_hotel_top_details_param", _paramHotelTopDetailsParam)
    return nil
}

// ParamHotelTopDetailsParam Getter
func (r AlitripHotelDetailInfoGetAPIRequest) GetParamHotelTopDetailsParam() *HotelTopDetailsParam {
    return r._paramHotelTopDetailsParam
}
