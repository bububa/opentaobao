package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeRefundOrderGet 零售通退款订单查询
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
func AlibabaLstTradeRefundOrderGet(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeRefundOrderGetAPIRequest, resp *lsttrade.AlibabaLstTradeRefundOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
