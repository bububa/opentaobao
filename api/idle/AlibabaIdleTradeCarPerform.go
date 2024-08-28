package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTradeCarPerform 二手车寄卖履约接口
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
func AlibabaIdleTradeCarPerform(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTradeCarPerformAPIRequest, resp *idle.AlibabaIdleTradeCarPerformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
