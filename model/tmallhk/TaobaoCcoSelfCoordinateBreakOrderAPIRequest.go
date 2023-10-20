package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoccoselfcoordinatebreakorderAPIRequest 天猫国际直购供应商毁单通知 API请求
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
type TaobaoccoselfcoordinatebreakorderAPIRequest struct {
	model.Params
	// 毁单请求参数
	_supplierBreakOrderRequest *SupplierBreakOrderRequest
}

// NewTaobaoccoselfcoordinatebreakorderRequest 初始化TaobaoccoselfcoordinatebreakorderAPIRequest对象
func NewTaobaoccoselfcoordinatebreakorderRequest() *TaobaoccoselfcoordinatebreakorderAPIRequest {
	return &TaobaoccoselfcoordinatebreakorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoccoselfcoordinatebreakorderAPIRequest) GetApiMethodName() string {
	return "taobao.cco.self.coordinate.break.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoccoselfcoordinatebreakorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoccoselfcoordinatebreakorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierBreakOrderRequest is SupplierBreakOrderRequest Setter
// 毁单请求参数
func (r *TaobaoccoselfcoordinatebreakorderAPIRequest) SetSupplierBreakOrderRequest(_supplierBreakOrderRequest *SupplierBreakOrderRequest) error {
	r._supplierBreakOrderRequest = _supplierBreakOrderRequest
	r.Set("supplier_break_order_request", _supplierBreakOrderRequest)
	return nil
}

// GetSupplierBreakOrderRequest SupplierBreakOrderRequest Getter
func (r TaobaoccoselfcoordinatebreakorderAPIRequest) GetSupplierBreakOrderRequest() *SupplierBreakOrderRequest {
	return r._supplierBreakOrderRequest
}
