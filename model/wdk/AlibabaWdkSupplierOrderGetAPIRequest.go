package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierOrderGetAPIRequest 五道口按订单号批量查询供应商正向订单 API请求
// alibaba.wdk.supplier.order.get
//
// 五道口按订单号批量查询供应商正向订单
type AlibabaWdkSupplierOrderGetAPIRequest struct {
	model.Params
	// 查询参数
	_supplierOrderQueryListRequest *SupplierOrderQueryListRequest
}

// NewAlibabaWdkSupplierOrderGetRequest 初始化AlibabaWdkSupplierOrderGetAPIRequest对象
func NewAlibabaWdkSupplierOrderGetRequest() *AlibabaWdkSupplierOrderGetAPIRequest {
	return &AlibabaWdkSupplierOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SupplierOrderQueryListRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierOrderGetAPIRequest) SetSupplierOrderQueryListRequest(_supplierOrderQueryListRequest *SupplierOrderQueryListRequest) error {
	r._supplierOrderQueryListRequest = _supplierOrderQueryListRequest
	r.Set("supplier_order_query_list_request", _supplierOrderQueryListRequest)
	return nil
}

// Get SupplierOrderQueryListRequest Getter
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
	return r._supplierOrderQueryListRequest
}
