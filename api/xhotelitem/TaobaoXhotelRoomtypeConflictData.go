package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
商家床型冲突数据接口 
taobao.xhotel.roomtype.conflict.data

商家床型冲突数据接口
*/
func TaobaoXhotelRoomtypeConflictData(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeConflictDataRequest, session string) (*xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
