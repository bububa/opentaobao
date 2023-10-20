package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSupplierOrderGetAPIRequest) Reset() {
	r._supplierOrderQueryListRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierOrderQueryListRequest is SupplierOrderQueryListRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierOrderGetAPIRequest) SetSupplierOrderQueryListRequest(_supplierOrderQueryListRequest *SupplierOrderQueryListRequest) error {
	r._supplierOrderQueryListRequest = _supplierOrderQueryListRequest
	r.Set("supplier_order_query_list_request", _supplierOrderQueryListRequest)
	return nil
}

// GetSupplierOrderQueryListRequest SupplierOrderQueryListRequest Getter
func (r AlibabaWdkSupplierOrderGetAPIRequest) GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
	return r._supplierOrderQueryListRequest
}

var poolAlibabaWdkSupplierOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSupplierOrderGetRequest()
	},
}

// GetAlibabaWdkSupplierOrderGetRequest 从 sync.Pool 获取 AlibabaWdkSupplierOrderGetAPIRequest
func GetAlibabaWdkSupplierOrderGetAPIRequest() *AlibabaWdkSupplierOrderGetAPIRequest {
	return poolAlibabaWdkSupplierOrderGetAPIRequest.Get().(*AlibabaWdkSupplierOrderGetAPIRequest)
}

// ReleaseAlibabaWdkSupplierOrderGetAPIRequest 将 AlibabaWdkSupplierOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSupplierOrderGetAPIRequest(v *AlibabaWdkSupplierOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkSupplierOrderGetAPIRequest.Put(v)
}
