package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoRoomtypeStatusUpdate top房型状态修改
// taobao.roomtype.status.update
//
// top房型状态修改
func TaobaoRoomtypeStatusUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoRoomtypeStatusUpdateAPIRequest, resp *xhotelitem.TaobaoRoomtypeStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
