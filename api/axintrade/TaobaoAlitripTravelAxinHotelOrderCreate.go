package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderCreate 酒店分销订单创建服务-阿信
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
func TaobaoAlitripTravelAxinHotelOrderCreate(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest, session string) (*axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse, error) {
	var resp axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
