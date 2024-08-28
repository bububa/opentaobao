package nlife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradePay 零售+平台支付订单
// alibaba.nlife.b2c.trade.pay
//
// 零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
func AlibabaNlifeB2cTradePay(ctx context.Context, clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradePayAPIRequest, resp *nlife.AlibabaNlifeB2cTradePayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
