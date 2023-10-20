package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalorderrefundAPIRequest 退款接口 API请求
// alibaba.alihealth.medical.order.refund
//
// 退款接口
type AlibabaalihealthmedicalorderrefundAPIRequest struct {
	model.Params
	// 入参
	_orderRefundRequest *OrderRefundRequest
}

// NewAlibabaalihealthmedicalorderrefundRequest 初始化AlibabaalihealthmedicalorderrefundAPIRequest对象
func NewAlibabaalihealthmedicalorderrefundRequest() *AlibabaalihealthmedicalorderrefundAPIRequest {
	return &AlibabaalihealthmedicalorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalorderrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRefundRequest is OrderRefundRequest Setter
// 入参
func (r *AlibabaalihealthmedicalorderrefundAPIRequest) SetOrderRefundRequest(_orderRefundRequest *OrderRefundRequest) error {
	r._orderRefundRequest = _orderRefundRequest
	r.Set("order_refund_request", _orderRefundRequest)
	return nil
}

// GetOrderRefundRequest OrderRefundRequest Getter
func (r AlibabaalihealthmedicalorderrefundAPIRequest) GetOrderRefundRequest() *OrderRefundRequest {
	return r._orderRefundRequest
}
