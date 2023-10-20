package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomsIncrement 房型库存推送接口（批量增量）
// taobao.xhotel.rooms.increment
//
// Room库存增量更新接口，用户仅需要更新ROOM库存中发生变化的库存日历即可。
func TaobaoXhotelRoomsIncrement(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomsIncrementAPIRequest, resp *xhotelitem.TaobaoXhotelRoomsIncrementAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
