package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeAdd 房型新增接口（ID重复变更新）
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
func TaobaoXhotelRoomtypeAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeAddAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
