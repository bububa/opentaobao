package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallhkclearanceinfosendAPIRequest 清关信息回调通知 API请求
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
type TmallhkclearanceinfosendAPIRequest struct {
	model.Params
	// 清关信息
	_orderClearanceInfoRequest *OrderClearanceInfoRequest
}

// NewTmallhkclearanceinfosendRequest 初始化TmallhkclearanceinfosendAPIRequest对象
func NewTmallhkclearanceinfosendRequest() *TmallhkclearanceinfosendAPIRequest {
	return &TmallhkclearanceinfosendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallhkclearanceinfosendAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.info.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallhkclearanceinfosendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallhkclearanceinfosendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderClearanceInfoRequest is OrderClearanceInfoRequest Setter
// 清关信息
func (r *TmallhkclearanceinfosendAPIRequest) SetOrderClearanceInfoRequest(_orderClearanceInfoRequest *OrderClearanceInfoRequest) error {
	r._orderClearanceInfoRequest = _orderClearanceInfoRequest
	r.Set("order_clearance_info_request", _orderClearanceInfoRequest)
	return nil
}

// GetOrderClearanceInfoRequest OrderClearanceInfoRequest Getter
func (r TmallhkclearanceinfosendAPIRequest) GetOrderClearanceInfoRequest() *OrderClearanceInfoRequest {
	return r._orderClearanceInfoRequest
}
