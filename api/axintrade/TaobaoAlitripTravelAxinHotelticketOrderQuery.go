package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderQuery 阿信度假交易订单查询接口
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
func TaobaoAlitripTravelAxinHotelticketOrderQuery(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
