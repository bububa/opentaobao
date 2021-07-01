package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeConfirmfeeGetAPIRequest
获取交易确认收货费用 API请求
taobao.trade.confirmfee.get

获取交易确认收货费用，可以获取主订单或子订单的确认收货费用 */
type TaobaoTradeConfirmfeeGetAPIRequest struct {
	model.Params
	// 交易主订单或子订单ID
	_tid int64
}

// New
