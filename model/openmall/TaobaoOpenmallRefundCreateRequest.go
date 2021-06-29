package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OpenMall退款单 APIRequest
taobao.openmall.refund.create

创建OpenMall退款单
如存在未完结的退款单，则返回该退款单ID
*/
type TaobaoOpenmallRefundCreateRequest struct {
    model.Params

    // 分销者联盟身份
    distributor   string 

    // 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
    goodsStatus   string 

    // 买家的退货描述
    refundDesc   string 

    // 退款金额，分
    refundFee   int64 

    // 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
    refundReason   string 

    // 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
    refundType   string 

    // 订单号
    tid   int64 

}

func NewTaobaoOpenmallRefundCreateRequest() *TaobaoOpenmallRefundCreateRequest{
    return &TaobaoOpenmallRefundCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundCreateRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.create"
}

func (r TaobaoOpenmallRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundCreateRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundCreateRequest) SetGoodsStatus(goodsStatus string) error {
    r.goodsStatus = goodsStatus
    r.Set("goods_status", goodsStatus)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetGoodsStatus() string {
    return r.goodsStatus
}

func (r *TaobaoOpenmallRefundCreateRequest) SetRefundDesc(refundDesc string) error {
    r.refundDesc = refundDesc
    r.Set("refund_desc", refundDesc)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetRefundDesc() string {
    return r.refundDesc
}

func (r *TaobaoOpenmallRefundCreateRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetRefundFee() int64 {
    return r.refundFee
}

func (r *TaobaoOpenmallRefundCreateRequest) SetRefundReason(refundReason string) error {
    r.refundReason = refundReason
    r.Set("refund_reason", refundReason)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetRefundReason() string {
    return r.refundReason
}

func (r *TaobaoOpenmallRefundCreateRequest) SetRefundType(refundType string) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetRefundType() string {
    return r.refundType
}

func (r *TaobaoOpenmallRefundCreateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenmallRefundCreateRequest) GetTid() int64 {
    return r.tid
}

