package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTravelTicketOrderRefundAPIRequest
飞猪门票退票结果通知 API请求
taobao.travel.ticket.order.refund

门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。 */
type TaobaoTravelTicketOrderRefundAPIRequest struct {
	model.Params
	// 下单时订单ID
	_orderId int64
	// 退票结果；1: 退票成功；2: 退票失败
	_refundStatus int64
	// 退票失败理由
	_refundFailureReason string
}

// NewTaobaoTravelTicketOrderRefundRequest 初始化TaobaoTravelTicketOrderRefundAPIRequest对象
func NewTaobaoTravelTicketOrderRefundRequest() *TaobaoTravelTicketOrderRefundAPIRequest {
	return &TaobaoTravelTicketOrderRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetApiMethodName() string {
	return "taobao.travel.ticket.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 下单时订单ID
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is RefundStatus Setter
// 退票结果；1: 退票成功；2: 退票失败
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundStatus(_refundStatus int64) error {
	r._refundStatus = _refundStatus
	r.Set("refund_status", _refundStatus)
	return nil
}

// Get RefundStatus Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundStatus() int64 {
	return r._refundStatus
}

// Set is RefundFailureReason Setter
// 退票失败理由
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundFailureReason(_refundFailureReason string) error {
	r._refundFailureReason = _refundFailureReason
	r.Set("refund_failure_reason", _refundFailureReason)
	return nil
}

// Get RefundFailureReason Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundFailureReason() string {
	return r._refundFailureReason
}
