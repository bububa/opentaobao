package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeUpdate 房型更新接口（ID不存在自动新增）
// taobao.xhotel.roomtype.update
//
// 酒店房型更新或添加
func TaobaoXhotelRoomtypeUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
