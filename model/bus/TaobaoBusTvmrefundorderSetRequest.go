package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机逆向退款接口 API请求
taobao.bus.tvmrefundorder.set

汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
*/
type TaobaoBusTvmrefundorderSetRequest struct {
    model.Params
    // 飞猪订单号
    _alitripOrderId   string
    // 分账退款明细
    _refundAccountInDetails   []RefundAccountInDetail
    // 退款金额（单位分） 票金额
    _refundAmount   int64
    // 退款原因
    _refundReason   string
    // 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
    _refundBatchNo   string
    // 保险退款详情
    _insuranceRefundDetails   []InsuranceRefundDetail
}

// 初始化TaobaoBusTvmrefundorderSetRequest对象
func NewTaobaoBusTvmrefundorderSetRequest() *TaobaoBusTvmrefundorderSetRequest{
    return &TaobaoBusTvmrefundorderSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmrefundorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmrefundorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmrefundorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmrefundorderSetRequest) SetAlitripOrderId(_alitripOrderId string) error {
    r._alitripOrderId = _alitripOrderId
    r.Set("alitrip_order_id", _alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetAlitripOrderId() string {
    return r._alitripOrderId
}
// RefundAccountInDetails Setter
// 分账退款明细
func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundAccountInDetails(_refundAccountInDetails []RefundAccountInDetail) error {
    r._refundAccountInDetails = _refundAccountInDetails
    r.Set("refund_account_in_details", _refundAccountInDetails)
    return nil
}

// RefundAccountInDetails Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetRefundAccountInDetails() []RefundAccountInDetail {
    return r._refundAccountInDetails
}
// RefundAmount Setter
// 退款金额（单位分） 票金额
func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundAmount(_refundAmount int64) error {
    r._refundAmount = _refundAmount
    r.Set("refund_amount", _refundAmount)
    return nil
}

// RefundAmount Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetRefundAmount() int64 {
    return r._refundAmount
}
// RefundReason Setter
// 退款原因
func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundReason(_refundReason string) error {
    r._refundReason = _refundReason
    r.Set("refund_reason", _refundReason)
    return nil
}

// RefundReason Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetRefundReason() string {
    return r._refundReason
}
// RefundBatchNo Setter
// 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundBatchNo(_refundBatchNo string) error {
    r._refundBatchNo = _refundBatchNo
    r.Set("refund_batch_no", _refundBatchNo)
    return nil
}

// RefundBatchNo Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetRefundBatchNo() string {
    return r._refundBatchNo
}
// InsuranceRefundDetails Setter
// 保险退款详情
func (r *TaobaoBusTvmrefundorderSetRequest) SetInsuranceRefundDetails(_insuranceRefundDetails []InsuranceRefundDetail) error {
    r._insuranceRefundDetails = _insuranceRefundDetails
    r.Set("insurance_refund_details", _insuranceRefundDetails)
    return nil
}

// InsuranceRefundDetails Getter
func (r TaobaoBusTvmrefundorderSetRequest) GetInsuranceRefundDetails() []InsuranceRefundDetail {
    return r._insuranceRefundDetails
}
