package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBaseinfoRoomGet 酒店房型与房价查询
// taobao.xhotel.baseinfo.room.get
//
// 根据outHid/hid获取酒店的房型和价格信息
func TaobaoXhotelBaseinfoRoomGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBaseinfoRoomGetAPIRequest, resp *xhotelitem.TaobaoXhotelBaseinfoRoomGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
