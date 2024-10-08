package nlife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradeCancel 零售+平台取消订单
// alibaba.nlife.b2c.trade.cancel
//
// 零售+平台取消订单接口
func AlibabaNlifeB2cTradeCancel(ctx context.Context, clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeCancelAPIRequest, resp *nlife.AlibabaNlifeB2cTradeCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
