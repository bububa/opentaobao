package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeOrderRefundListQueryAPIRequest
查询退款单列表(卖家视角) API请求
alibaba.lst.trade.order.refund.list.query

查询退款单列表(卖家视角) */
type AlibabaLstTradeOrderRefundListQueryAPIRequest struct {
	model.Params
	// 输入参数
	_param *TopLstSupplierOrderRefundQuery
}

// New
