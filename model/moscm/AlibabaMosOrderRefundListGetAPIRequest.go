package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOrderRefundListGetAPIRequest
批量查询交易退货信息 API请求
alibaba.mos.order.refund.list.get

批量查询多个退货单的退货明细 */
type AlibabaMosOrderRefundListGetAPIRequest struct {
	model.Params
	// 退换货查询条件
	_rmaCriteria *RmaCriteria
	// 分页信息
	_paginator *Paginator
}

// New
