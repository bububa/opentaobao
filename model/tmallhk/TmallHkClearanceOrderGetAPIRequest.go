package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearanceordergetAPIRequest 天猫国际订单清关信息 API请求
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
type TmallhkclearanceordergetAPIRequest struct {
	model.Params
	// 入参封装类型
	_clearanceOrderRequest *ClearanceOrderRequest
}

// NewTmallhkclearanceordergetRequest 初始化TmallhkclearanceordergetAPIRequest对象
func NewTmallhkclearanceordergetRequest() *TmallhkclearanceordergetAPIRequest {
	return &TmallhkclearanceordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallhkclearanceordergetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallhkclearanceordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallhkclearanceordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClearanceOrderRequest is ClearanceOrderRequest Setter
// 入参封装类型
func (r *TmallhkclearanceordergetAPIRequest) SetClearanceOrderRequest(_clearanceOrderRequest *ClearanceOrderRequest) error {
	r._clearanceOrderRequest = _clearanceOrderRequest
	r.Set("clearance_order_request", _clearanceOrderRequest)
	return nil
}

// GetClearanceOrderRequest ClearanceOrderRequest Getter
func (r TmallhkclearanceordergetAPIRequest) GetClearanceOrderRequest() *ClearanceOrderRequest {
	return r._clearanceOrderRequest
}
