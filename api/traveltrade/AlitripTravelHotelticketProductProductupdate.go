package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketProductProductupdate 产品批量变更通知
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
func AlitripTravelHotelticketProductProductupdate(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketProductProductupdateAPIRequest, resp *traveltrade.AlitripTravelHotelticketProductProductupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
