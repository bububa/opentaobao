package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeGet 房型查询接口
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
func TaobaoXhotelRoomtypeGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeGetAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
