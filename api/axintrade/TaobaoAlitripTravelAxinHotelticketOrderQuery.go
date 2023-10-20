package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelticketOrderQuery 阿信度假交易订单查询接口
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
func TaobaoAlitripTravelAxinHotelticketOrderQuery(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelticketOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
