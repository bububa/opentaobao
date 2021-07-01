package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoRefundRefundactionsGetAPIRequest
判断该订单能执行的逆向操作集合列表 API请求
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表 */
type CainiaoRefundRefundactionsGetAPIRequest struct {
	model.Params
	// 子订单ID
	_orderId string
}

// New
