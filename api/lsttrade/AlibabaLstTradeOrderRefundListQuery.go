package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeOrderRefundListQuery 查询退款单列表(卖家视角)
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
func AlibabaLstTradeOrderRefundListQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeOrderRefundListQueryAPIRequest, resp *lsttrade.AlibabaLstTradeOrderRefundListQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
