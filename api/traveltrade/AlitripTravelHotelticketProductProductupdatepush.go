package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketProductProductupdatepush 产品批量变更推送通知
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
func AlitripTravelHotelticketProductProductupdatepush(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketProductProductupdatepushAPIRequest, resp *traveltrade.AlitripTravelHotelticketProductProductupdatepushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
