package car

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarOrderComplete 服务完成API
// taobao.alitrip.car.order.complete
//
// 用来接收服务商订单流程完成信息
func TaobaoAlitripCarOrderComplete(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderCompleteAPIRequest, resp *car.TaobaoAlitripCarOrderCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
