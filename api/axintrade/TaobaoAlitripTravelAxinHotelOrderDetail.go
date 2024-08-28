package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderDetail 阿信酒店分销-订单详情接口
// taobao.alitrip.travel.axin.hotel.order.detail
//
// 阿信酒店订单详情的读取接口
func TaobaoAlitripTravelAxinHotelOrderDetail(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderDetailAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
