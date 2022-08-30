package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelRoomMatch 阿信酒店房型匹配
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
func TaobaoAlitripTravelAxinHotelRoomMatch(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
