package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoRoomtypeStatusUpdate top房型状态修改
// taobao.roomtype.status.update
//
// top房型状态修改
func TaobaoRoomtypeStatusUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoRoomtypeStatusUpdateAPIRequest, session string) (*xhotelitem.TaobaoRoomtypeStatusUpdateAPIResponse, error) {
	var resp xhotelitem.TaobaoRoomtypeStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
