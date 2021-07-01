package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeInvoiceAmountGetAPIRequest
获取订单应开票金额 API请求
taobao.trade.invoice.amount.get

订单应开票金额计算 */
type TaobaoTradeInvoiceAmountGetAPIRequest struct {
	model.Params
	// 业务订单ID
	_tid int64
}

// New
