package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// CainiaoRefundRefundactionsDisplay 退货退款操作的展示信息(展现给买家)
// cainiao.refund.refundactions.display
//
// 退货退款操作的展示信息(展现给买家)
func CainiaoRefundRefundactionsDisplay(ctx context.Context, clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsDisplayAPIRequest, resp *trade.CainiaoRefundRefundactionsDisplayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
