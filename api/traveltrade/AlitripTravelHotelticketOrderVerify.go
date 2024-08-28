package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderVerify 订单核销通知
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
func AlitripTravelHotelticketOrderVerify(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderVerifyAPIRequest, resp *traveltrade.AlitripTravelHotelticketOrderVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
