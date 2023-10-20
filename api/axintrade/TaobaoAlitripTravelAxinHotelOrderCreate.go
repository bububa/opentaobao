package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripTravelAxinHotelOrderCreate 酒店分销订单创建服务-阿信
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
func TaobaoAlitripTravelAxinHotelOrderCreate(clt *core.SDKClient, req *axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest, resp *axintrade.TaobaoAlitripTravelAxinHotelOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
