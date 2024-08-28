package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractSensorTradeBuy 手淘下单能力开放
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
func AlibabaInteractSensorTradeBuy(ctx context.Context, clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorTradeBuyAPIRequest, resp *mtopopen.AlibabaInteractSensorTradeBuyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
