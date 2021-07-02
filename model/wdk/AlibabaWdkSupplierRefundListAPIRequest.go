package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierRefundListAPIRequest 五道口按供应商拉取退款单 API请求
// alibaba.wdk.supplier.refund.list
//
// 五道口按供应商拉取退款单
type AlibabaWdkSupplierRefundListAPIRequest struct {
	model.Params
	// 查询参数
	_supplierRefundQueryRequest *SupplierRefundQueryRequest
}

// NewAlibabaWdkSupplierRefundListRequest 初始化AlibabaWdkSupplierRefundListAPIRequest对象
func NewAlibabaWdkSupplierRefundListRequest() *AlibabaWdkSupplierRefundListAPIRequest {
	return &AlibabaWdkSupplierRefundListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierRefundQueryRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierRefundListAPIRequest) SetSupplierRefundQueryRequest(_supplierRefundQueryRequest *SupplierRefundQueryRequest) error {
	r._supplierRefundQueryRequest = _supplierRefundQueryRequest
	r.Set("supplier_refund_query_request", _supplierRefundQueryRequest)
	return nil
}

// Get SupplierRefundQueryRequest Getter
func (r AlibabaWdkSupplierRefundListAPIRequest) GetSupplierRefundQueryRequest() *SupplierRefundQueryRequest {
	return r._supplierRefundQueryRequest
}
