package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderVerify 订单核销通知
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
func AlitripTravelHotelticketOrderVerify(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderVerifyAPIRequest, resp *traveltrade.AlitripTravelHotelticketOrderVerifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
