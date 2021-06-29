package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据HID获取所有卖家房型匹配关系 API请求
alitrip.hotel.hstdf.shotel.roomtype.mappings.list

根据HID获取所有卖家房型匹配关系
*/
type AlitripHotelHstdfShotelRoomtypeMappingsListRequest struct {
    model.Params
    // HID
    _hid   int64
}

// 初始化AlitripHotelHstdfShotelRoomtypeMappingsListRequest对象
func NewAlitripHotelHstdfShotelRoomtypeMappingsListRequest() *AlitripHotelHstdfShotelRoomtypeMappingsListRequest{
    return &AlitripHotelHstdfShotelRoomtypeMappingsListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.roomtype.mappings.list"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// HID
func (r *AlitripHotelHstdfShotelRoomtypeMappingsListRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r AlitripHotelHstdfShotelRoomtypeMappingsListRequest) GetHid() int64 {
    return r._hid
}
