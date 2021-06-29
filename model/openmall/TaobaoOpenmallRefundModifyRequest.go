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
    _refundType   string
    // 退款金额，分
    _refundFee   int64
    // 买家的退货描述
    _refundDesc   string
    // 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
    _goodsStatus   string
    // 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
    _refundReason   string
    // 分销者联盟身份
    _distributor   string
    // 退款单ID
    _refundId   int64
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
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundType(_refundType string) error {
    r._refundType = _refundType
    r.Set("refund_type", _refundType)
    return nil
}

// RefundType Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundType() string {
    return r._refundType
}
// RefundFee Setter
// 退款金额，分
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundFee(_refundFee int64) error {
    r._refundFee = _refundFee
    r.Set("refund_fee", _refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundFee() int64 {
    return r._refundFee
}
// RefundDesc Setter
// 买家的退货描述
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundDesc(_refundDesc string) error {
    r._refundDesc = _refundDesc
    r.Set("refund_desc", _refundDesc)
    return nil
}

// RefundDesc Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundDesc() string {
    return r._refundDesc
}
// GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoOpenmallRefundModifyRequest) SetGoodsStatus(_goodsStatus string) error {
    r._goodsStatus = _goodsStatus
    r.Set("goods_status", _goodsStatus)
    return nil
}

// GoodsStatus Getter
func (r TaobaoOpenmallRefundModifyRequest) GetGoodsStatus() string {
    return r._goodsStatus
}
// RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundReason(_refundReason string) error {
    r._refundReason = _refundReason
    r.Set("refund_reason", _refundReason)
    return nil
}

// RefundReason Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundReason() string {
    return r._refundReason
}
// Distributor Setter
// 分销者联盟身份
func (r *TaobaoOpenmallRefundModifyRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundModifyRequest) GetDistributor() string {
    return r._distributor
}
// RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundModifyRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundModifyRequest) GetRefundId() int64 {
    return r._refundId
}
