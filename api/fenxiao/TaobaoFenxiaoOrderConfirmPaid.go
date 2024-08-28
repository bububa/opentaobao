package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoOrderConfirmPaid 确认收款
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
func TaobaoFenxiaoOrderConfirmPaid(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIRequest, resp *fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
