package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderVerify 订单核销通知
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
func AlitripTravelHotelticketOrderVerify(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderVerifyAPIRequest, session string) (*traveltrade.AlitripTravelHotelticketOrderVerifyAPIResponse, error) {
	var resp traveltrade.AlitripTravelHotelticketOrderVerifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
