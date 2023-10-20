package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketRefundOrderrefund 阿信度假业务申请退款
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
func TaobaoAlitripTravelAxinHotelticketRefundOrderrefund(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
