package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomGet room实体查询
// taobao.xhotel.room.get
//
// 此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
func TaobaoXhotelRoomGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomGetAPIRequest, resp *xhotelitem.TaobaoXhotelRoomGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
