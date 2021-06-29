package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页静态信息 API请求
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
type AlitripHotelDetailStaticinfoGetRequest struct {
    model.Params
    // 酒店详情页静态信息参数
    _paramHotelTopStaticDetailsParam   *HotelTopStaticDetailsParam
}

// 初始化AlitripHotelDetailStaticinfoGetRequest对象
func NewAlitripHotelDetailStaticinfoGetRequest() *AlitripHotelDetailStaticinfoGetRequest{
    return &AlitripHotelDetailStaticinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelDetailStaticinfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.detail.staticinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelDetailStaticinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamHotelTopStaticDetailsParam Setter
// 酒店详情页静态信息参数
func (r *AlitripHotelDetailStaticinfoGetRequest) SetParamHotelTopStaticDetailsParam(_paramHotelTopStaticDetailsParam *HotelTopStaticDetailsParam) error {
    r._paramHotelTopStaticDetailsParam = _paramHotelTopStaticDetailsParam
    r.Set("param_hotel_top_static_details_param", _paramHotelTopStaticDetailsParam)
    return nil
}

// ParamHotelTopStaticDetailsParam Getter
func (r AlitripHotelDetailStaticinfoGetRequest) GetParamHotelTopStaticDetailsParam() *HotelTopStaticDetailsParam {
    return r._paramHotelTopStaticDetailsParam
}
