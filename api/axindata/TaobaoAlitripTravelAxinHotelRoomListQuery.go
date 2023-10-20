package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelRoomListQuery 阿信酒店分销-标准酒店房型列表查询
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
func TaobaoAlitripTravelAxinHotelRoomListQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelRoomListQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
