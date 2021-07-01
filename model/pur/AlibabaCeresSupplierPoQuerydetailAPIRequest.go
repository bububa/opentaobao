package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCeresSupplierPoQuerydetailAPIRequest
采购供应商订单明细查询接口 API请求
alibaba.ceres.supplier.po.querydetail

采购供应商订单明细查询接口 */
type AlibabaCeresSupplierPoQuerydetailAPIRequest struct {
	model.Params
	// 订单编号
	_poNo string
}

// New
