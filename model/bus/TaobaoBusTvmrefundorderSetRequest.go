package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机逆向退款接口 APIRequest
taobao.bus.tvmrefundorder.set

汽车票线下自助机 逆向退票接口；用于已出票完成后，再发起退款（注意这是售后退款，如出票异常但是告诉我们出票成功，后续给客户退款，需要调用这个接口，一般开放给财务。出票过程中的失败，请直接调用出票接口并且传递false标志，我们会自动退款。）
*/
type TaobaoBusTvmrefundorderSetRequest struct {
    model.Params

    // 飞猪订单号
    alitripOrderId   string 

    // 分账退款明细
    refundAccountInDetails   []RefundAccountInDetail 

    // 退款金额（单位分） 票金额
    refundAmount   int64 

    // 退款原因
    refundReason   string 

    // 批次号必须唯一，同一批次号只能退款一次 （多账号分润的该值 填写refundAccountInDetails中批次号的任意一个即可
    refundBatchNo   string 

    // 保险退款详情
    insuranceRefundDetails   []InsuranceRefundDetail 

}

func NewTaobaoBusTvmrefundorderSetRequest() *TaobaoBusTvmrefundorderSetRequest{
    return &TaobaoBusTvmrefundorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmrefundorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmrefundorder.set"
}

func (r TaobaoBusTvmrefundorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmrefundorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundAccountInDetails(refundAccountInDetails []RefundAccountInDetail) error {
    r.refundAccountInDetails = refundAccountInDetails
    r.Set("refund_account_in_details", refundAccountInDetails)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetRefundAccountInDetails() []RefundAccountInDetail {
    return r.refundAccountInDetails
}

func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundAmount(refundAmount int64) error {
    r.refundAmount = refundAmount
    r.Set("refund_amount", refundAmount)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetRefundAmount() int64 {
    return r.refundAmount
}

func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetRefundReason() string {
    return r.refundReason
}

func (r *TaobaoBusTvmrefundorderSetRequest) SetRefundBatchNo(refundBatchNo string) error {
    r.refundBatchNo = refundBatchNo
    r.Set("refund_batch_no", refundBatchNo)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetRefundBatchNo() string {
    return r.refundBatchNo
}

func (r *TaobaoBusTvmrefundorderSetRequest) SetInsuranceRefundDetails(insuranceRefundDetails []InsuranceRefundDetail) error {
    r.insuranceRefundDetails = insuranceRefundDetails
    r.Set("insurance_refund_details", insuranceRefundDetails)
    return nil
}

func (r TaobaoBusTvmrefundorderSetRequest) GetInsuranceRefundDetails() []InsuranceRefundDetail {
    return r.insuranceRefundDetails
}

