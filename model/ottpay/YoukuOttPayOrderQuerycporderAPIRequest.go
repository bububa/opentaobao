package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQuerycporderAPIRequest
查询支付订单对应cp订单号 API请求
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号 */
type YoukuOttPayOrderQuerycporderAPIRequest struct {
	model.Params
	// 支付对应订单
	_gatewayOrder string
}

// New
