package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个卖家房型下的所有标准房型 API请求
alitrip.hotel.hstdf.shotel.exportsroomtype

导出一个卖家酒店下的所有标准房型
*/
type AlitripHotelHstdfShotelExportsroomtypeRequest struct {
    model.Params
    // 卖家酒店id
    _hid   int64
}

// 初始化AlitripHotelHstdfShotelExportsroomtypeRequest对象
func NewAlitripHotelHstdfShotelExportsroomtypeRequest() *AlitripHotelHstdfShotelExportsroomtypeRequest{
    return &AlitripHotelHstdfShotelExportsroomtypeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.exportsroomtype"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 卖家酒店id
func (r *AlitripHotelHstdfShotelExportsroomtypeRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r AlitripHotelHstdfShotelExportsroomtypeRequest) GetHid() int64 {
    return r._hid
}
