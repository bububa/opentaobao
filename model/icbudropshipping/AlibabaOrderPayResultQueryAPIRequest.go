package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOrderPayResultQueryAPIRequest
alibaba查询订单支付结果 API请求
alibaba.order.pay.result.query

alibaba查询订单支付结果 */
type AlibabaOrderPayResultQueryAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// New
