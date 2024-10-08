package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderPay 阿信酒店分销订单支付
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
func TaobaoAlitripTravelAxinHotelOrderPay(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderPayAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderPayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
