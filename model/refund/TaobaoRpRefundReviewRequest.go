package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
审核退款单 API请求
taobao.rp.refund.review

审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
*/
type TaobaoRpRefundReviewRequest struct {
    model.Params
    // 退款单编号
    refundId   int64
    // 审核人姓名
    operator   string
    // 退款阶段，可选值：售中：onsale，售后：aftersale
    refundPhase   string
    // 退款最后更新时间，以时间戳的方式表示
    refundVersion   int64
    // 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
    result   bool
    // 审核留言
    message   string
}

// 初始化TaobaoRpRefundReviewRequest对象
func NewTaobaoRpRefundReviewRequest() *TaobaoRpRefundReviewRequest{
    return &TaobaoRpRefundReviewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpRefundReviewRequest) GetApiMethodName() string {
    return "taobao.rp.refund.review"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpRefundReviewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单编号
func (r *TaobaoRpRefundReviewRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpRefundReviewRequest) GetRefundId() int64 {
    return r.refundId
}
// Operator Setter
// 审核人姓名
func (r *TaobaoRpRefundReviewRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoRpRefundReviewRequest) GetOperator() string {
    return r.operator
}
// RefundPhase Setter
// 退款阶段，可选值：售中：onsale，售后：aftersale
func (r *TaobaoRpRefundReviewRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpRefundReviewRequest) GetRefundPhase() string {
    return r.refundPhase
}
// RefundVersion Setter
// 退款最后更新时间，以时间戳的方式表示
func (r *TaobaoRpRefundReviewRequest) SetRefundVersion(refundVersion int64) error {
    r.refundVersion = refundVersion
    r.Set("refund_version", refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpRefundReviewRequest) GetRefundVersion() int64 {
    return r.refundVersion
}
// Result Setter
// 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核）
func (r *TaobaoRpRefundReviewRequest) SetResult(result bool) error {
    r.result = result
    r.Set("result", result)
    return nil
}

// Result Getter
func (r TaobaoRpRefundReviewRequest) GetResult() bool {
    return r.result
}
// Message Setter
// 审核留言
func (r *TaobaoRpRefundReviewRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

// Message Getter
func (r TaobaoRpRefundReviewRequest) GetMessage() string {
    return r.message
}
