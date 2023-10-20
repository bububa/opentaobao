package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderRefund 退款结结果通知
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
func AlitripTravelHotelticketOrderRefund(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderRefundAPIRequest, session string) (*traveltrade.AlitripTravelHotelticketOrderRefundAPIResponse, error) {
	var resp traveltrade.AlitripTravelHotelticketOrderRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
