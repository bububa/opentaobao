package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketProductDetail 阿信酒景套餐产品详情查询
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
func TaobaoAlitripTravelAxinHotelticketProductDetail(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
