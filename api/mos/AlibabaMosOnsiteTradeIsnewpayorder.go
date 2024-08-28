package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosOnsiteTradeIsnewpayorder 是否为新支付订单
// alibaba.mos.onsite.trade.isnewpayorder
//
// 退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
func AlibabaMosOnsiteTradeIsnewpayorder(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradeIsnewpayorderAPIRequest, resp *mos.AlibabaMosOnsiteTradeIsnewpayorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
