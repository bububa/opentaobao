package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdksupplierordergetAPIRequest 五道口按订单号批量查询供应商正向订单 API请求
// alibaba.wdk.supplier.order.get
//
// 五道口按订单号批量查询供应商正向订单
type AlibabawdksupplierordergetAPIRequest struct {
	model.Params
	// 查询参数
	_supplierOrderQueryListRequest *SupplierOrderQueryListRequest
}

// NewAlibabawdksupplierordergetRequest 初始化AlibabawdksupplierordergetAPIRequest对象
func NewAlibabawdksupplierordergetRequest() *AlibabawdksupplierordergetAPIRequest {
	return &AlibabawdksupplierordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdksupplierordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.supplier.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdksupplierordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdksupplierordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierOrderQueryListRequest is SupplierOrderQueryListRequest Setter
// 查询参数
func (r *AlibabawdksupplierordergetAPIRequest) SetSupplierOrderQueryListRequest(_supplierOrderQueryListRequest *SupplierOrderQueryListRequest) error {
	r._supplierOrderQueryListRequest = _supplierOrderQueryListRequest
	r.Set("supplier_order_query_list_request", _supplierOrderQueryListRequest)
	return nil
}

// GetSupplierOrderQueryListRequest SupplierOrderQueryListRequest Getter
func (r AlibabawdksupplierordergetAPIRequest) GetSupplierOrderQueryListRequest() *SupplierOrderQueryListRequest {
	return r._supplierOrderQueryListRequest
}
