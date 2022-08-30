package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderRefund 阿信酒店分销订单退款API
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
func TaobaoAlitripTravelAxinHotelOrderRefund(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelOrderRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
