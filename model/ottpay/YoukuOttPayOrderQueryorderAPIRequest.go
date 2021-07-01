package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQueryorderAPIRequest
查询订单 API请求
youku.ott.pay.order.queryorder

通过订单号查询订单信息 */
type YoukuOttPayOrderQueryorderAPIRequest struct {
	model.Params
	// 订单号
	_orderNo string
}

// New
