package hotelhstdf

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelhstdf"
)

/* 
根据HID获取所有卖家房型匹配关系 
alitrip.hotel.hstdf.shotel.roomtype.mappings.list

根据HID获取所有卖家房型匹配关系
*/
func AlitripHotelHstdfShotelRoomtypeMappingsList(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIRequest, session string) (*hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse, error) {
    var resp hotelhstdf.AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
