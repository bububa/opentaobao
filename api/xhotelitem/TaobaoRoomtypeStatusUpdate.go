package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoRoomtypeStatusUpdate top房型状态修改
// taobao.roomtype.status.update
//
// top房型状态修改
func TaobaoRoomtypeStatusUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoRoomtypeStatusUpdateAPIRequest, resp *xhotelitem.TaobaoRoomtypeStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
