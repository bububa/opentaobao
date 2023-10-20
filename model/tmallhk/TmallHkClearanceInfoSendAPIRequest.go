package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceInfoSendAPIRequest 清关信息回调通知 API请求
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
type TmallHkClearanceInfoSendAPIRequest struct {
	model.Params
	// 清关信息
	_orderClearanceInfoRequest *OrderClearanceInfoRequest
}

// NewTmallHkClearanceInfoSendRequest 初始化TmallHkClearanceInfoSendAPIRequest对象
func NewTmallHkClearanceInfoSendRequest() *TmallHkClearanceInfoSendAPIRequest {
	return &TmallHkClearanceInfoSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkClearanceInfoSendAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.info.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkClearanceInfoSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallHkClearanceInfoSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderClearanceInfoRequest is OrderClearanceInfoRequest Setter
// 清关信息
func (r *TmallHkClearanceInfoSendAPIRequest) SetOrderClearanceInfoRequest(_orderClearanceInfoRequest *OrderClearanceInfoRequest) error {
	r._orderClearanceInfoRequest = _orderClearanceInfoRequest
	r.Set("order_clearance_info_request", _orderClearanceInfoRequest)
	return nil
}

// GetOrderClearanceInfoRequest OrderClearanceInfoRequest Getter
func (r TmallHkClearanceInfoSendAPIRequest) GetOrderClearanceInfoRequest() *OrderClearanceInfoRequest {
	return r._orderClearanceInfoRequest
}
