package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketRefundOrderrefund 阿信度假业务申请退款
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
func TaobaoAlitripTravelAxinHotelticketRefundOrderrefund(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
