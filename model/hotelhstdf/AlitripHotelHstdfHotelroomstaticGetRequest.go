package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据类型查询静态字段 APIRequest
alitrip.hotel.hstdf.hotelroomstatic.get

根据类型查询分页静态字段
*/
type AlitripHotelHstdfHotelroomstaticGetRequest struct {
    model.Params

    // 参数封装
    paramGetHotelRoomStaticParam   *GetHotelRoomStaticParam 

}

func NewAlitripHotelHstdfHotelroomstaticGetRequest() *AlitripHotelHstdfHotelroomstaticGetRequest{
    return &AlitripHotelHstdfHotelroomstaticGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.hotelroomstatic.get"
}

func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfHotelroomstaticGetRequest) SetParamGetHotelRoomStaticParam(paramGetHotelRoomStaticParam *GetHotelRoomStaticParam) error {
    r.paramGetHotelRoomStaticParam = paramGetHotelRoomStaticParam
    r.Set("param_get_hotel_room_static_param", paramGetHotelRoomStaticParam)
    return nil
}

func (r AlitripHotelHstdfHotelroomstaticGetRequest) GetParamGetHotelRoomStaticParam() *GetHotelRoomStaticParam {
    return r.paramGetHotelRoomStaticParam
}

