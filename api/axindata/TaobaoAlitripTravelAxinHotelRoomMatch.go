package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelRoomMatch 阿信酒店房型匹配
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
func TaobaoAlitripTravelAxinHotelRoomMatch(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelRoomMatchAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelRoomMatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
