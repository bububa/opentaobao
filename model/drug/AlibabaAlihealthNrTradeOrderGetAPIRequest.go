package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrTradeOrderGetAPIRequest
获取订单详情 API请求
alibaba.alihealth.nr.trade.order.get

阿里健康O2O，获取订单详情 */
type AlibabaAlihealthNrTradeOrderGetAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// New
