package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

// AlibabaLstTradeSellerOrderDetailQuery 订单详情查看(卖家视角)
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
func AlibabaLstTradeSellerOrderDetailQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeSellerOrderDetailQueryAPIRequest, resp *lsttrade.AlibabaLstTradeSellerOrderDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
