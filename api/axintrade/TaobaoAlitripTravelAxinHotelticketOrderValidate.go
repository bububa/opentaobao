package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderValidate 阿信度假业务交易试单接口
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
func TaobaoAlitripTravelAxinHotelticketOrderValidate(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelticketOrderValidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
