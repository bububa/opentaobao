package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarOrderAccept 确认接单
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
func TaobaoAlitripCarOrderAccept(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderAcceptAPIRequest, resp *car.TaobaoAlitripCarOrderAcceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
