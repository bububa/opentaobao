package lsttrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsttrade"
)

/* AlibabaLstTradeSellerOrderListQuery
订单列表查看(卖家视角)
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表 */
func AlibabaLstTradeSellerOrderListQuery(clt *core.SDKClient, req *lsttrade.AlibabaLstTradeSellerOrderListQueryAPIRequest, session string) (*lsttrade.AlibabaLstTradeSellerOrderListQueryAPIResponse, error) {
	var resp lsttrade.AlibabaLstTradeSellerOrderListQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
