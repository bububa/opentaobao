package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOrderListGetAPIRequest
批量查询订单交易 API请求
alibaba.mos.order.list.get

批量查询交易信息 */
type AlibabaMosOrderListGetAPIRequest struct {
	model.Params
	// 订单查询条件
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// New
