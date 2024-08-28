package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomUpdate 房型库存推送接口（全量推送）
// taobao.xhotel.room.update
//
// 此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
func TaobaoXhotelRoomUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRoomUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
