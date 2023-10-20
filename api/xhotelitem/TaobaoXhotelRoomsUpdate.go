package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelroomsupdate 房型共享库存推送接口（批量全量）
// taobao.xhotel.rooms.update
//
// 此接口用于更新多个集市酒店商品房态信息，根据传入的gids更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。是全量更新，非增量，会把之前的房态进行覆盖。
func Taobaoxhotelroomsupdate(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelroomsupdateAPIRequest, session string) (*xhotelitem.TaobaoxhotelroomsupdateAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelroomsupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
