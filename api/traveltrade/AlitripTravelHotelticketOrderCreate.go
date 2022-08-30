package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderCreate 创单(支付订单)通知
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
func AlitripTravelHotelticketOrderCreate(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderCreateAPIRequest, session string) (*traveltrade.AlitripTravelHotelticketOrderCreateAPIResponse, error) {
	var resp traveltrade.AlitripTravelHotelticketOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
