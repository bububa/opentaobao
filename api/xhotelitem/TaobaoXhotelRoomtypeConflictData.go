package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeConflictData 商家床型冲突数据接口
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
func TaobaoXhotelRoomtypeConflictData(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeConflictDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
