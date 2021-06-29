package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
匹配标准房型以及卖家房型 API请求
alitrip.hotel.hstdf.shotel.matchsroomself

匹配卖家房型以及标准房型，返回匹配结果
*/
type AlitripHotelHstdfShotelMatchsroomselfRequest struct {
    model.Params
    // SroomTypeMatchParam
    param0   *SroomTypeMatchParam
}

// 初始化AlitripHotelHstdfShotelMatchsroomselfRequest对象
func NewAlitripHotelHstdfShotelMatchsroomselfRequest() *AlitripHotelHstdfShotelMatchsroomselfRequest{
    return &AlitripHotelHstdfShotelMatchsroomselfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.matchsroomself"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// SroomTypeMatchParam
func (r *AlitripHotelHstdfShotelMatchsroomselfRequest) SetParam0(param0 *SroomTypeMatchParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlitripHotelHstdfShotelMatchsroomselfRequest) GetParam0() *SroomTypeMatchParam {
    return r.param0
}
