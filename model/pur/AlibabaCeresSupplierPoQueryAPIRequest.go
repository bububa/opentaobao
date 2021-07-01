package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCeresSupplierPoQueryAPIRequest
采购供应商订单查询接口 API请求
alibaba.ceres.supplier.po.query

采购供应商订单查询接口 */
type AlibabaCeresSupplierPoQueryAPIRequest struct {
	model.Params
	// 订单创建日期开始时间
	_startDate string
	// 订单创建日期结束时间
	_endDate string
	// 订单状态
	_status string
}

// New
