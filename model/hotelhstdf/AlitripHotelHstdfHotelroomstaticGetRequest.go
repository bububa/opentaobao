package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据类型查询静态字段 API请求
alitrip.hotel.hstdf.hotelroomstatic.get

根据类型查询分页静态字段
*/
type AlitripHotelHstdfHotelroomstaticGetRequest struct {
    model.Params
    // 参数封装
    paramGetHotelRoomStaticParam   *GetHotelRoomStaticParam
}

// 初始化AlitripHotelHstdfHotelroomstaticGetRequest对象
func NewAlitripHotelHstdfHotelroomstaticGetRequest() *AlitripHotelHstdfHotelroomstaticGetRequest{
    return &AlitripHotelHstdfHotelroomstaticGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.hotelroomstatic.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamGetHotelRoomStaticParam Setter
// 参数封装
func (r *AlitripHotelHstdfHotelroomstaticGetRequest) SetParamGetHotelRoomStaticParam(paramGetHotelRoomStaticParam *GetHotelRoomStaticParam) error {
    r.paramGetHotelRoomStaticParam = paramGetHotelRoomStaticParam
    r.Set("param_get_hotel_room_static_param", paramGetHotelRoomStaticParam)
    return nil
}

// ParamGetHotelRoomStaticParam Getter
func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetParamGetHotelRoomStaticParam() *GetHotelRoomStaticParam {
    return r.paramGetHotelRoomStaticParam
}
