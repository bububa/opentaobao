package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个卖家房型下的所有标准房型 APIRequest
alitrip.hotel.hstdf.shotel.exportsroomtype

导出一个卖家酒店下的所有标准房型
*/
type AlitripHotelHstdfShotelExportsroomtypeRequest struct {
    model.Params

    // 卖家酒店id
    hid   int64 

}

func NewAlitripHotelHstdfShotelExportsroomtypeRequest() *AlitripHotelHstdfShotelExportsroomtypeRequest{
    return &AlitripHotelHstdfShotelExportsroomtypeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.exportsroomtype"
}

func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelExportsroomtypeRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetHid() int64 {
    return r.hid
}

