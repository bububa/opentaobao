package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderRefund 阿信酒店分销订单退款API
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
func TaobaoAlitripTravelAxinHotelOrderRefund(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
