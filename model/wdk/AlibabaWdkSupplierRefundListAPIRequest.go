package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksupplierrefundlistAPIRequest 五道口按供应商拉取退款单 API请求
// alibaba.wdk.supplier.refund.list
//
// 五道口按供应商拉取退款单
type AlibabawdksupplierrefundlistAPIRequest struct {
	model.Params
	// 查询参数
	_supplierRefundQueryRequest *SupplierRefundQueryRequest
}

// NewAlibabawdksupplierrefundlistRequest 初始化AlibabawdksupplierrefundlistAPIRequest对象
func NewAlibabawdksupplierrefundlistRequest() *AlibabawdksupplierrefundlistAPIRequest {
	return &AlibabawdksupplierrefundlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdksupplierrefundlistAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdksupplierrefundlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdksupplierrefundlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierRefundQueryRequest is SupplierRefundQueryRequest Setter
// 查询参数
func (r *AlibabawdksupplierrefundlistAPIRequest) SetSupplierRefundQueryRequest(_supplierRefundQueryRequest *SupplierRefundQueryRequest) error {
	r._supplierRefundQueryRequest = _supplierRefundQueryRequest
	r.Set("supplier_refund_query_request", _supplierRefundQueryRequest)
	return nil
}

// GetSupplierRefundQueryRequest SupplierRefundQueryRequest Getter
func (r AlibabawdksupplierrefundlistAPIRequest) GetSupplierRefundQueryRequest() *SupplierRefundQueryRequest {
	return r._supplierRefundQueryRequest
}
