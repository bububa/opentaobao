package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
酒店房型与房价查询 
taobao.xhotel.baseinfo.room.get

根据outHid/hid获取酒店的房型和价格信息
*/
func TaobaoXhotelBaseinfoRoomGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBaseinfoRoomGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelBaseinfoRoomGetAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelBaseinfoRoomGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
