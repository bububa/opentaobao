package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketProductProductupdatepush 产品批量变更推送通知
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
func AlitripTravelHotelticketProductProductupdatepush(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketProductProductupdatepushAPIRequest, resp *traveltrade.AlitripTravelHotelticketProductProductupdatepushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
