package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeDeletePublic 商家删除房型数据接口
// taobao.xhotel.roomtype.delete.public
//
// 房型删除TOP接口
func TaobaoXhotelRoomtypeDeletePublic(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeDeletePublicAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeDeletePublicAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
