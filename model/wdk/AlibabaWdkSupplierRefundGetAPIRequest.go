package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierRefundGetAPIRequest 五道口按订单号批量查询供应商退款单 API请求
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
type AlibabaWdkSupplierRefundGetAPIRequest struct {
	model.Params
	// 查询入参
	_supplierRefundQueryListRequest *SupplierRefundQueryListRequest
}

// NewAlibabaWdkSupplierRefundGetRequest 初始化AlibabaWdkSupplierRefundGetAPIRequest对象
func NewAlibabaWdkSupplierRefundGetRequest() *AlibabaWdkSupplierRefundGetAPIRequest {
	return &AlibabaWdkSupplierRefundGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSupplierRefundGetAPIRequest) Reset() {
	r._supplierRefundQueryListRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierRefundQueryListRequest is SupplierRefundQueryListRequest Setter
// 查询入参
func (r *AlibabaWdkSupplierRefundGetAPIRequest) SetSupplierRefundQueryListRequest(_supplierRefundQueryListRequest *SupplierRefundQueryListRequest) error {
	r._supplierRefundQueryListRequest = _supplierRefundQueryListRequest
	r.Set("supplier_refund_query_list_request", _supplierRefundQueryListRequest)
	return nil
}

// GetSupplierRefundQueryListRequest SupplierRefundQueryListRequest Getter
func (r AlibabaWdkSupplierRefundGetAPIRequest) GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
	return r._supplierRefundQueryListRequest
}

var poolAlibabaWdkSupplierRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSupplierRefundGetRequest()
	},
}

// GetAlibabaWdkSupplierRefundGetRequest 从 sync.Pool 获取 AlibabaWdkSupplierRefundGetAPIRequest
func GetAlibabaWdkSupplierRefundGetAPIRequest() *AlibabaWdkSupplierRefundGetAPIRequest {
	return poolAlibabaWdkSupplierRefundGetAPIRequest.Get().(*AlibabaWdkSupplierRefundGetAPIRequest)
}

// ReleaseAlibabaWdkSupplierRefundGetAPIRequest 将 AlibabaWdkSupplierRefundGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSupplierRefundGetAPIRequest(v *AlibabaWdkSupplierRefundGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkSupplierRefundGetAPIRequest.Put(v)
}
