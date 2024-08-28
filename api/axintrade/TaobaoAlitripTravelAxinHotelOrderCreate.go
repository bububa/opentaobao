package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderCreate 酒店分销订单创建服务-阿信
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
func TaobaoAlitripTravelAxinHotelOrderCreate(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
