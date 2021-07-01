package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailCommissionOrderQueryAPIRequest
分销订单查询 API请求
alibaba.retail.commission.order.query

查询商家的分销订单 */
type AlibabaRetailCommissionOrderQueryAPIRequest struct {
	model.Params
	// 页码，默认第一页
	_pageNo int64
	// 页大小，默认每页十条
	_pageSize int64
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	_endPayTime string
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	_startPayTime string
}

// New
