package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketProductDetail 阿信酒景套餐产品详情查询
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
func TaobaoAlitripTravelAxinHotelticketProductDetail(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelticketProductDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
