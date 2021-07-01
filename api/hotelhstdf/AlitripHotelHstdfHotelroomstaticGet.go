package hotelhstdf

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelhstdf"
)

/* 
根据类型查询静态字段 
alitrip.hotel.hstdf.hotelroomstatic.get

根据类型查询分页静态字段
*/
func AlitripHotelHstdfHotelroomstaticGet(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfHotelroomstaticGetAPIRequest, session string) (*hotelhstdf.AlitripHotelHstdfHotelroomstaticGetAPIResponse, error) {
    var resp hotelhstdf.AlitripHotelHstdfHotelroomstaticGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
