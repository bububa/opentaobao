package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据HID获取所有卖家房型匹配关系 APIRequest
alitrip.hotel.hstdf.shotel.roomtype.mappings.list

根据HID获取所有卖家房型匹配关系
*/
type AlitripHotelHstdfShotelRoomtypeMappingsListRequest struct {
    model.Params

    // HID
    hid   int64 

}

func NewAlitripHotelHstdfShotelRoomtypeMappingsListRequest() *AlitripHotelHstdfShotelRoomtypeMappingsListRequest{
    return &AlitripHotelHstdfShotelRoomtypeMappingsListRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.roomtype.mappings.list"
}

func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelRoomtypeMappingsListRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetHid() int64 {
    return r.hid
}

