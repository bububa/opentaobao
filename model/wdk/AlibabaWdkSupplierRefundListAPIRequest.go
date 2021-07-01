package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSupplierRefundListAPIRequest
五道口按供应商拉取退款单 API请求
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单 */
type AlibabaWdkSupplierRefundListAPIRequest struct {
	model.Params
	// 查询参数
	_supplierRefundQueryRequest *SupplierRefundQueryRequest
}

// New
