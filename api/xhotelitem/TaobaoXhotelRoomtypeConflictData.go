package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeConflictData 商家床型冲突数据接口
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
func TaobaoXhotelRoomtypeConflictData(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
