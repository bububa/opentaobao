package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪门票退票结果通知 API请求
taobao.travel.ticket.order.refund

门票系统商通过TOP接口通知飞猪门票是否退票成功，以及退票数量。
*/
type TaobaoTravelTicketOrderRefundRequest struct {
    model.Params
    // 下单时订单ID
    orderId   int64
    // 退票结果；1: 退票成功；2: 退票失败
    refundStatus   int64
    // 退票失败理由
    refundFailureReason   string
}

// 初始化TaobaoTravelTicketOrderRefundRequest对象
func NewTaobaoTravelTicketOrderRefundRequest() *TaobaoTravelTicketOrderRefundRequest{
    return &TaobaoTravelTicketOrderRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTravelTicketOrderRefundRequest) GetApiMethodName() string {
    return "taobao.travel.ticket.order.refund"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTravelTicketOrderRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 下单时订单ID
func (r *TaobaoTravelTicketOrderRefundRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTravelTicketOrderRefundRequest) GetOrderId() int64 {
    return r.orderId
}
// RefundStatus Setter
// 退票结果；1: 退票成功；2: 退票失败
func (r *TaobaoTravelTicketOrderRefundRequest) SetRefundStatus(refundStatus int64) error {
    r.refundStatus = refundStatus
    r.Set("refund_status", refundStatus)
    return nil
}

// RefundStatus Getter
func (r TaobaoTravelTicketOrderRefundRequest) GetRefundStatus() int64 {
    return r.refundStatus
}
// RefundFailureReason Setter
// 退票失败理由
func (r *TaobaoTravelTicketOrderRefundRequest) SetRefundFailureReason(refundFailureReason string) error {
    r.refundFailureReason = refundFailureReason
    r.Set("refund_failure_reason", refundFailureReason)
    return nil
}

// RefundFailureReason Getter
func (r TaobaoTravelTicketOrderRefundRequest) GetRefundFailureReason() string {
    return r.refundFailureReason
}
