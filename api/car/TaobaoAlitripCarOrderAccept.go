package car

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarOrderAccept 确认接单
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
func TaobaoAlitripCarOrderAccept(ctx context.Context, clt *core.SDKClient, req *car.TaobaoAlitripCarOrderAcceptAPIRequest, resp *car.TaobaoAlitripCarOrderAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
