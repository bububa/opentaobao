package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeRefundOrderGetAPIRequest
零售通退款订单查询 API请求
alibaba.lst.trade.refund.order.get

零售通退款订单查询 */
type AlibabaLstTradeRefundOrderGetAPIRequest struct {
	model.Params
	// 退款单id
	_refundId string
	// 主订单id
	_mainOrderId int64
}

// New
