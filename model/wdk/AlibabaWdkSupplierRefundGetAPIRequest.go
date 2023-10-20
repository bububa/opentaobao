package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksupplierrefundgetAPIRequest 五道口按订单号批量查询供应商退款单 API请求
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
type AlibabawdksupplierrefundgetAPIRequest struct {
	model.Params
	// 查询入参
	_supplierRefundQueryListRequest *SupplierRefundQueryListRequest
}

// NewAlibabawdksupplierrefundgetRequest 初始化AlibabawdksupplierrefundgetAPIRequest对象
func NewAlibabawdksupplierrefundgetRequest() *AlibabawdksupplierrefundgetAPIRequest {
	return &AlibabawdksupplierrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdksupplierrefundgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdksupplierrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdksupplierrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierRefundQueryListRequest is SupplierRefundQueryListRequest Setter
// 查询入参
func (r *AlibabawdksupplierrefundgetAPIRequest) SetSupplierRefundQueryListRequest(_supplierRefundQueryListRequest *SupplierRefundQueryListRequest) error {
	r._supplierRefundQueryListRequest = _supplierRefundQueryListRequest
	r.Set("supplier_refund_query_list_request", _supplierRefundQueryListRequest)
	return nil
}

// GetSupplierRefundQueryListRequest SupplierRefundQueryListRequest Getter
func (r AlibabawdksupplierrefundgetAPIRequest) GetSupplierRefundQueryListRequest() *SupplierRefundQueryListRequest {
	return r._supplierRefundQueryListRequest
}
