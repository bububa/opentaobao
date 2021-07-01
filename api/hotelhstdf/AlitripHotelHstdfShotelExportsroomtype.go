package hotelhstdf

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelhstdf"
)

/* 
导出一个卖家房型下的所有标准房型 
alitrip.hotel.hstdf.shotel.exportsroomtype

导出一个卖家酒店下的所有标准房型
*/
func AlitripHotelHstdfShotelExportsroomtype(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIRequest, session string) (*hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIResponse, error) {
    var resp hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
