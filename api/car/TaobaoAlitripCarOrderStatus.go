package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarOrderStatus 商家订单状态改变通知接口（神州专车接口）
// taobao.alitrip.car.order.status
//
// 商家订单状态改变通知接口，神州专车专用接口！
func TaobaoAlitripCarOrderStatus(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderStatusAPIRequest, resp *car.TaobaoAlitripCarOrderStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
