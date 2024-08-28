package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// CainiaoRefundRefundactionsJudgement 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
// cainiao.refund.refundactions.judgement
//
// 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
func CainiaoRefundRefundactionsJudgement(ctx context.Context, clt *core.SDKClient, req *trade.CainiaoRefundRefundactionsJudgementAPIRequest, resp *trade.CainiaoRefundRefundactionsJudgementAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
