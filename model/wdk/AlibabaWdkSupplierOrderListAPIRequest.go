package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSupplierOrderListAPIRequest
五道口供应商维度正向订单拉取 API请求
alibaba.wdk.supplier.order.list

五道口供应商维度正向订单拉取 */
type AlibabaWdkSupplierOrderListAPIRequest struct {
	model.Params
	// 查询参数
	_supplierOrderQueryRequest *SupplierOrderQueryRequest
}

// New
