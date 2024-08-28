package alitripcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// TaobaoAlitripCarOrderAgentCancel 司机或客服取消订单
// taobao.alitrip.car.order.agent.cancel
//
// 司机或客服取消订单后通知飞猪订单取消
func TaobaoAlitripCarOrderAgentCancel(ctx context.Context, clt *core.SDKClient, req *alitripcar.TaobaoAlitripCarOrderAgentCancelAPIRequest, resp *alitripcar.TaobaoAlitripCarOrderAgentCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
