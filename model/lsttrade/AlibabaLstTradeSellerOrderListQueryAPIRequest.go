package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOrderListQueryAPIRequest
订单列表查看(卖家视角) API请求
alibaba.lst.trade.seller.order.list.query

卖家视角订单查询，查询授权经销商订单列表 */
type AlibabaLstTradeSellerOrderListQueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// New
