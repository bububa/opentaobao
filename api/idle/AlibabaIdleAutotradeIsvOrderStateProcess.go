package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAutotradeIsvOrderStateProcess 闲鱼订单状态推进
// alibaba.idle.autotrade.isv.order.state.process
//
// 闲鱼订单状态推进
func AlibabaIdleAutotradeIsvOrderStateProcess(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleAutotradeIsvOrderStateProcessAPIRequest, resp *idle.AlibabaIdleAutotradeIsvOrderStateProcessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
