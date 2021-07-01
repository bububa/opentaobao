package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOrderDetailQueryAPIRequest
订单详情查看(卖家视角) API请求
alibaba.lst.trade.seller.order.detail.query

订单详情查看(卖家视角) */
type AlibabaLstTradeSellerOrderDetailQueryAPIRequest struct {
	model.Params
	// 入参
	_param *LstTradeGetSellerOrderListParam
}

// New
