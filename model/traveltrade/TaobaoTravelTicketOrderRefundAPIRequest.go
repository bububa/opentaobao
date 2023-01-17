package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTravelTicketOrderRefundAPIRequest 飞猪门票退票结果通知 API请求
// taobao.travel.ticket.order.refund
//
// 门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。
type TaobaoTravelTicketOrderRefundAPIRequest struct {
	model.Params
	// 退票失败理由
	_refundFailureReason string
	// 退票的批次号
	_refundBatchNo string
	// 退票数量
	_refundNum int64
	// 下单时订单ID
	_orderId int64
	// 退票结果；1: 退票成功；2: 退票失败
	_refundStatus int64
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
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundFailureReason is RefundFailureReason Setter
// 退票失败理由
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundFailureReason(_refundFailureReason string) error {
	r._refundFailureReason = _refundFailureReason
	r.Set("refund_failure_reason", _refundFailureReason)
	return nil
}

// GetRefundFailureReason RefundFailureReason Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundFailureReason() string {
	return r._refundFailureReason
}

// SetRefundBatchNo is RefundBatchNo Setter
// 退票的批次号
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundBatchNo(_refundBatchNo string) error {
	r._refundBatchNo = _refundBatchNo
	r.Set("refund_batch_no", _refundBatchNo)
	return nil
}

// GetRefundBatchNo RefundBatchNo Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundBatchNo() string {
	return r._refundBatchNo
}

// SetRefundNum is RefundNum Setter
// 退票数量
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundNum(_refundNum int64) error {
	r._refundNum = _refundNum
	r.Set("refund_num", _refundNum)
	return nil
}

// GetRefundNum RefundNum Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundNum() int64 {
	return r._refundNum
}

// SetOrderId is OrderId Setter
// 下单时订单ID
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetRefundStatus is RefundStatus Setter
// 退票结果；1: 退票成功；2: 退票失败
func (r *TaobaoTravelTicketOrderRefundAPIRequest) SetRefundStatus(_refundStatus int64) error {
	r._refundStatus = _refundStatus
	r.Set("refund_status", _refundStatus)
	return nil
}

// GetRefundStatus RefundStatus Getter
func (r TaobaoTravelTicketOrderRefundAPIRequest) GetRefundStatus() int64 {
	return r._refundStatus
}
