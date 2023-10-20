package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelHouseRoomtypeAdd 民宿房型信息添加
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
func TaobaoXhotelHouseRoomtypeAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelHouseRoomtypeAddAPIRequest, resp *xhotelitem.TaobaoXhotelHouseRoomtypeAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
