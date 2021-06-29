package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪门票退票结果通知 APIRequest
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

func NewTaobaoTravelTicketOrderRefundRequest() *TaobaoTravelTicketOrderRefundRequest{
    return &TaobaoTravelTicketOrderRefundRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTravelTicketOrderRefundRequest) GetApiMethodName() string {
    return "taobao.travel.ticket.order.refund"
}

func (r TaobaoTravelTicketOrderRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTravelTicketOrderRefundRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoTravelTicketOrderRefundRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoTravelTicketOrderRefundRequest) SetRefundStatus(refundStatus int64) error {
    r.refundStatus = refundStatus
    r.Set("refund_status", refundStatus)
    return nil
}

func (r TaobaoTravelTicketOrderRefundRequest) GetRefundStatus() int64 {
    return r.refundStatus
}

func (r *TaobaoTravelTicketOrderRefundRequest) SetRefundFailureReason(refundFailureReason string) error {
    r.refundFailureReason = refundFailureReason
    r.Set("refund_failure_reason", refundFailureReason)
    return nil
}

func (r TaobaoTravelTicketOrderRefundRequest) GetRefundFailureReason() string {
    return r.refundFailureReason
}

