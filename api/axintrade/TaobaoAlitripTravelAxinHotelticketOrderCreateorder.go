package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderCreateorder 阿信度假业务创单并支付接口
// taobao.alitrip.travel.axin.hotelticket.order.createorder
//
// 阿信度假业务创单并支付接口
func TaobaoAlitripTravelAxinHotelticketOrderCreateorder(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
