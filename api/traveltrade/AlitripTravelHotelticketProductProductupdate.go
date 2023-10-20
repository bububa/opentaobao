package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketProductProductupdate 产品批量变更通知
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
func AlitripTravelHotelticketProductProductupdate(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketProductProductupdateAPIRequest, session string) (*traveltrade.AlitripTravelHotelticketProductProductupdateAPIResponse, error) {
	var resp traveltrade.AlitripTravelHotelticketProductProductupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
