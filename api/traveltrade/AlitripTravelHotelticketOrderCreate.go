package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderCreate 创单(支付订单)通知
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
func AlitripTravelHotelticketOrderCreate(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderCreateAPIRequest, resp *traveltrade.AlitripTravelHotelticketOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
