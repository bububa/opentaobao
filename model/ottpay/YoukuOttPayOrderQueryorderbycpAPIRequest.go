package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQueryorderbycpAPIRequest
订单查询接口(cp订单号查询) API请求
youku.ott.pay.order.queryorderbycp

给商户服务端查询订单状态 */
type YoukuOttPayOrderQueryorderbycpAPIRequest struct {
	model.Params
	// cp订单号
	_cpOrderNo string
}

// New
