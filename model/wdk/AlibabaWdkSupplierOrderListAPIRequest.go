package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierOrderListAPIRequest 五道口供应商维度正向订单拉取 API请求
// alibaba.wdk.supplier.order.list
//
// 五道口供应商维度正向订单拉取
type AlibabaWdkSupplierOrderListAPIRequest struct {
	model.Params
	// 查询参数
	_supplierOrderQueryRequest *SupplierOrderQueryRequest
}

// NewAlibabaWdkSupplierOrderListRequest 初始化AlibabaWdkSupplierOrderListAPIRequest对象
func NewAlibabaWdkSupplierOrderListRequest() *AlibabaWdkSupplierOrderListAPIRequest {
	return &AlibabaWdkSupplierOrderListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSupplierOrderListAPIRequest) Reset() {
	r._supplierOrderQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierOrderListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.order.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierOrderListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSupplierOrderListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierOrderQueryRequest is SupplierOrderQueryRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierOrderListAPIRequest) SetSupplierOrderQueryRequest(_supplierOrderQueryRequest *SupplierOrderQueryRequest) error {
	r._supplierOrderQueryRequest = _supplierOrderQueryRequest
	r.Set("supplier_order_query_request", _supplierOrderQueryRequest)
	return nil
}

// GetSupplierOrderQueryRequest SupplierOrderQueryRequest Getter
func (r AlibabaWdkSupplierOrderListAPIRequest) GetSupplierOrderQueryRequest() *SupplierOrderQueryRequest {
	return r._supplierOrderQueryRequest
}

var poolAlibabaWdkSupplierOrderListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSupplierOrderListRequest()
	},
}

// GetAlibabaWdkSupplierOrderListRequest 从 sync.Pool 获取 AlibabaWdkSupplierOrderListAPIRequest
func GetAlibabaWdkSupplierOrderListAPIRequest() *AlibabaWdkSupplierOrderListAPIRequest {
	return poolAlibabaWdkSupplierOrderListAPIRequest.Get().(*AlibabaWdkSupplierOrderListAPIRequest)
}

// ReleaseAlibabaWdkSupplierOrderListAPIRequest 将 AlibabaWdkSupplierOrderListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSupplierOrderListAPIRequest(v *AlibabaWdkSupplierOrderListAPIRequest) {
	v.Reset()
	poolAlibabaWdkSupplierOrderListAPIRequest.Put(v)
}
