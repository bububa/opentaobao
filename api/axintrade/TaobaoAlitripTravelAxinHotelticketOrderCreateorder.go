package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderCreateorder 阿信度假业务创单并支付接口
// taobao.alitrip.travel.axin.hotelticket.order.createorder
//
// 阿信度假业务创单并支付接口
func TaobaoAlitripTravelAxinHotelticketOrderCreateorder(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
