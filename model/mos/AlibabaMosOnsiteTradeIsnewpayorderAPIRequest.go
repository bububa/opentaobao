package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradeIsnewpayorderAPIRequest
是否为新支付订单 API请求
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单 */
type AlibabaMosOnsiteTradeIsnewpayorderAPIRequest struct {
	model.Params
	// 外部订单号
	_outTradeNo string
}

// New
