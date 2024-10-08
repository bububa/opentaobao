package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeInvoiceAmountGet 获取订单应开票金额
// taobao.trade.invoice.amount.get
//
// 订单应开票金额计算
func TaobaoTradeInvoiceAmountGet(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTradeInvoiceAmountGetAPIRequest, resp *tbtrade.TaobaoTradeInvoiceAmountGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
