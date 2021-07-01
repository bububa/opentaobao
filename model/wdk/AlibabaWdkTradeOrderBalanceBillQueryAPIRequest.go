package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeOrderBalanceBillQueryAPIRequest
分页拉取订单数据 API请求
alibaba.wdk.trade.order.balance.bill.query

提供接口供外部调用，分页拉取订单数据 */
type AlibabaWdkTradeOrderBalanceBillQueryAPIRequest struct {
	model.Params
	// 入参
	_orderBalanceBillRequest *OrderBalanceBillRequest
}

// New
