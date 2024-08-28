package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// CainiaoRefundRefundactionsGet 判断该订单能执行的逆向操作集合列表
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
func CainiaoRefundRefundactionsGet(ctx context.Context, clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsGetAPIRequest, resp *trade.CainiaoRefundRefundactionsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
