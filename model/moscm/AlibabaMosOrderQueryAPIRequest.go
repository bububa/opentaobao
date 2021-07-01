package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOrderQueryAPIRequest
批量查询订单信息 API请求
alibaba.mos.order.query

查询多笔交易信息 */
type AlibabaMosOrderQueryAPIRequest struct {
	model.Params
	// 订单查询
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// New
