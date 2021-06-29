package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改OpenMall退款申请 APIRequest
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

func NewTaobaoOpenmallRefundModifyRequest() *TaobaoOpenmallRefundModifyRequest{
    return &TaobaoOpenmallRefundModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundModifyRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.modify"
}

func (r TaobaoOpenmallRefundModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundModifyRequest) SetRefundType(refundType string) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetRefundType() string {
    return r.refundType
}

func (r *TaobaoOpenmallRefundModifyRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetRefundFee() int64 {
    return r.refundFee
}

func (r *TaobaoOpenmallRefundModifyRequest) SetRefundDesc(refundDesc string) error {
    r.refundDesc = refundDesc
    r.Set("refund_desc", refundDesc)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetRefundDesc() string {
    return r.refundDesc
}

func (r *TaobaoOpenmallRefundModifyRequest) SetGoodsStatus(goodsStatus string) error {
    r.goodsStatus = goodsStatus
    r.Set("goods_status", goodsStatus)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetGoodsStatus() string {
    return r.goodsStatus
}

func (r *TaobaoOpenmallRefundModifyRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetRefundReason() string {
    return r.refundReason
}

func (r *TaobaoOpenmallRefundModifyRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundModifyRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundModifyRequest) GetRefundId() int64 {
    return r.refundId
}

