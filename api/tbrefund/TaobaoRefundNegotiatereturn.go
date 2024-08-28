package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundNegotiatereturn 协商退货退款
// taobao.refund.negotiatereturn
//
// 协商退货退款
func TaobaoRefundNegotiatereturn(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundNegotiatereturnAPIRequest, resp *tbrefund.TaobaoRefundNegotiatereturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
