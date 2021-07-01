package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeReceivetimeDelayAPIRequest
延长交易收货时间 API请求
taobao.trade.receivetime.delay

延长交易收货时间 */
type TaobaoTradeReceivetimeDelayAPIRequest struct {
	model.Params
	// 主订单号
	_tid int64
	// 延长收货的天数，可选值为：3, 5, 7, 10。
	_days int64
}

// New
