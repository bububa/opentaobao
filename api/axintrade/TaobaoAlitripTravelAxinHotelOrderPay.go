package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderPay 阿信酒店分销订单支付
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
func TaobaoAlitripTravelAxinHotelOrderPay(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderPayAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelOrderPayAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelOrderPayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
