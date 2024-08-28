package lsttrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeOrderRefundListQuery 查询退款单列表(卖家视角)
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
func AlibabaLstTradeOrderRefundListQuery(ctx context.Context, clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderRefundListQueryAPIRequest, resp *lsttrade.AlibabaLstTradeOrderRefundListQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
