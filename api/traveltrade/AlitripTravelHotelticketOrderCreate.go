package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderCreate 创单(支付订单)通知
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
func AlitripTravelHotelticketOrderCreate(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderCreateAPIRequest, resp *traveltrade.AlitripTravelHotelticketOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
