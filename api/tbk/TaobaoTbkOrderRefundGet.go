package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkOrderRefundGet 淘宝客-推广者-全量维权退款订单查询
// taobao.tbk.order.refund.get
//
// 淘宝客账户下全量维权退款订单查询
func TaobaoTbkOrderRefundGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkOrderRefundGetAPIRequest, resp *tbk.TaobaoTbkOrderRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
