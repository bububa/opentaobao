package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketProductList 阿信酒景套餐产品列表查询
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
func TaobaoAlitripTravelAxinHotelticketProductList(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketProductListAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketProductListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
