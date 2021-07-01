package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderQueryauthstateAPIRequest
查询连包签约状态 API请求
youku.ott.pay.order.queryauthstate

查询CP用户连包商品签约状态 */
type YoukuOttPayOrderQueryauthstateAPIRequest struct {
	model.Params
	// 原始签约订单号
	_originalCpOrderNo string
}

// New
