package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
匹配标准房型以及卖家房型 APIRequest
alitrip.hotel.hstdf.shotel.matchsroomself

匹配卖家房型以及标准房型，返回匹配结果
*/
type AlitripHotelHstdfShotelMatchsroomselfRequest struct {
    model.Params

    // SroomTypeMatchParam
    param0   *SroomTypeMatchParam 

}

func NewAlitripHotelHstdfShotelMatchsroomselfRequest() *AlitripHotelHstdfShotelMatchsroomselfRequest{
    return &AlitripHotelHstdfShotelMatchsroomselfRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.matchsroomself"
}

func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelMatchsroomselfRequest) SetParam0(param0 *SroomTypeMatchParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetParam0() *SroomTypeMatchParam {
    return r.param0
}

