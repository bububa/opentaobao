package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSupplierRefundGetAPIRequest
五道口按订单号批量查询供应商退款单 API请求
alibaba.wdk.supplier.refund.get

五道口按订单号批量查询供应商退款单 */
type AlibabaWdkSupplierRefundGetAPIRequest struct {
	model.Params
	// 查询入参
	_supplierRefundQueryListRequest *SupplierRefundQueryListRequest
}

// New
