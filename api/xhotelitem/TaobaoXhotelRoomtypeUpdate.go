package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeUpdate 房型更新接口（ID不存在自动新增）
// taobao.xhotel.roomtype.update
//
// 酒店房型更新或添加
func TaobaoXhotelRoomtypeUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
