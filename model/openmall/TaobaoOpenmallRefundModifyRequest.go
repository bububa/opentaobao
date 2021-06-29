package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改OpenMall退款申请 API请求
taobao.openmall.refund.modify

修改OpenMall退款申请
*/
type TaobaoOpenmallRefundModifyRequest struct {
    model.Params
    // 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
    refundType   string
    // 退款金额，分
    refundFee   int64
    // 买家的退货描述
    refundDesc   string
    // 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
    goodsStatus   string
    // 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
    refundReason   string
    // 分销者联盟身份
    distributor   string
    // 退款单ID
    refundId   int64
}

// 初始化TaobaoOpenmallRefundModifyRequest对象
func NewTaobaoOpenmallRefundModifyRequest() *TaobaoOpenmallRefundModifyRequest{
    return &TaobaoOpenmallRefundModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundModifyRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundType Setter
// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundType(refundType string) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

// RefundType Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundType() string {
    return r.refundType
}
// RefundFee Setter
// 退款金额，分
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundFee() int64 {
    return r.refundFee
}
// RefundDesc Setter
// 买家的退货描述
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundDesc(refundDesc string) error {
    r.refundDesc = refundDesc
    r.Set("refund_desc", refundDesc)
    return nil
}

// RefundDesc Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundDesc() string {
    return r.refundDesc
}
// GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoOpenmallRefundModifyRequest) SetGoodsStatus(goodsStatus string) error {
    r.goodsStatus = goodsStatus
    r.Set("goods_status", goodsStatus)
    return nil
}

// GoodsStatus Getter
func (r TaobaoOpenmallRefundModifyRequest) GetGoodsStatus() string {
    return r.goodsStatus
}
// RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

// RefundReason Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundReason() string {
    return r.refundReason
}
// Distributor Setter
// 分销者联盟身份
func (r *TaobaoOpenmallRefundModifyRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundModifyRequest) GetDistributor() string {
    return r.distributor
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundId() int64 {
    return r.refundId
}
