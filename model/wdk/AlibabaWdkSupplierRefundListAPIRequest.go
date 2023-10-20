package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSupplierRefundListAPIRequest) Reset() {
	r._supplierRefundQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSupplierRefundListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSupplierRefundListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSupplierRefundListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierRefundQueryRequest is SupplierRefundQueryRequest Setter
// 查询参数
func (r *AlibabaWdkSupplierRefundListAPIRequest) SetSupplierRefundQueryRequest(_supplierRefundQueryRequest *SupplierRefundQueryRequest) error {
	r._supplierRefundQueryRequest = _supplierRefundQueryRequest
	r.Set("supplier_refund_query_request", _supplierRefundQueryRequest)
	return nil
}

// GetSupplierRefundQueryRequest SupplierRefundQueryRequest Getter
func (r AlibabaWdkSupplierRefundListAPIRequest) GetSupplierRefundQueryRequest() *SupplierRefundQueryRequest {
	return r._supplierRefundQueryRequest
}

var poolAlibabaWdkSupplierRefundListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSupplierRefundListRequest()
	},
}

// GetAlibabaWdkSupplierRefundListRequest 从 sync.Pool 获取 AlibabaWdkSupplierRefundListAPIRequest
func GetAlibabaWdkSupplierRefundListAPIRequest() *AlibabaWdkSupplierRefundListAPIRequest {
	return poolAlibabaWdkSupplierRefundListAPIRequest.Get().(*AlibabaWdkSupplierRefundListAPIRequest)
}

// ReleaseAlibabaWdkSupplierRefundListAPIRequest 将 AlibabaWdkSupplierRefundListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSupplierRefundListAPIRequest(v *AlibabaWdkSupplierRefundListAPIRequest) {
	v.Reset()
	poolAlibabaWdkSupplierRefundListAPIRequest.Put(v)
}
