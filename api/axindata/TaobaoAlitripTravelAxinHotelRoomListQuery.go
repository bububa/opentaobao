package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelRoomListQuery 阿信酒店分销-标准酒店房型列表查询
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
func TaobaoAlitripTravelAxinHotelRoomListQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
