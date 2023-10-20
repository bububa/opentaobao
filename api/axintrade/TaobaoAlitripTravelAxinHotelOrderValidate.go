package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderValidate 阿信酒店订单数据校验
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
func TaobaoAlitripTravelAxinHotelOrderValidate(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelOrderValidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
