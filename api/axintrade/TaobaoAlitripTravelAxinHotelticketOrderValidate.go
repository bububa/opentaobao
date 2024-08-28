package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderValidate 阿信度假业务交易试单接口
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
func TaobaoAlitripTravelAxinHotelticketOrderValidate(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
