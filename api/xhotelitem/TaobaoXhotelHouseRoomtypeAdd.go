package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelHouseRoomtypeAdd 民宿房型信息添加
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
func TaobaoXhotelHouseRoomtypeAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelHouseRoomtypeAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelHouseRoomtypeAddAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelHouseRoomtypeAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
