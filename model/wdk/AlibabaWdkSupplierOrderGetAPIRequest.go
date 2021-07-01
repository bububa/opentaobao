package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSupplierOrderGetAPIRequest
五道口按订单号批量查询供应商正向订单 API请求
alibaba.wdk.supplier.order.get

五道口按订单号批量查询供应商正向订单 */
type AlibabaWdkSupplierOrderGetAPIRequest struct {
	model.Params
	// 查询参数
	_supplierOrderQueryListRequest *SupplierOrderQueryListRequest
}

// New
