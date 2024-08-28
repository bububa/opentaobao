package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderValidate 阿信酒店订单数据校验
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
func TaobaoAlitripTravelAxinHotelOrderValidate(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
