package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OpenMall退款单 API请求
taobao.openmall.refund.create

创建OpenMall退款单
如存在未完结的退款单，则返回该退款单ID
*/
type TaobaoOpenmallRefundCreateRequest struct {
    model.Params
    // 分销者联盟身份
    _distributor   string
    // 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
    _goodsStatus   string
    // 买家的退货描述
    _refundDesc   string
    // 退款金额，分
    _refundFee   int64
    // 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
    _refundReason   string
    // 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
    _refundType   string
    // 订单号
    _tid   int64
}

// 初始化TaobaoOpenmallRefundCreateRequest对象
func NewTaobaoOpenmallRefundCreateRequest() *TaobaoOpenmallRefundCreateRequest{
    return &TaobaoOpenmallRefundCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundCreateRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者联盟身份
func (r *TaobaoOpenmallRefundCreateRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundCreateRequest) GetDistributor() string {
    return r._distributor
}
// GoodsStatus Setter
// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
func (r *TaobaoOpenmallRefundCreateRequest) SetGoodsStatus(_goodsStatus string) error {
    r._goodsStatus = _goodsStatus
    r.Set("goods_status", _goodsStatus)
    return nil
}

// GoodsStatus Getter
func (r TaobaoOpenmallRefundCreateRequest) GetGoodsStatus() string {
    return r._goodsStatus
}
// RefundDesc Setter
// 买家的退货描述
func (r *TaobaoOpenmallRefundCreateRequest) SetRefundDesc(_refundDesc string) error {
    r._refundDesc = _refundDesc
    r.Set("refund_desc", _refundDesc)
    return nil
}

// RefundDesc Getter
func (r TaobaoOpenmallRefundCreateRequest) GetRefundDesc() string {
    return r._refundDesc
}
// RefundFee Setter
// 退款金额，分
func (r *TaobaoOpenmallRefundCreateRequest) SetRefundFee(_refundFee int64) error {
    r._refundFee = _refundFee
    r.Set("refund_fee", _refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoOpenmallRefundCreateRequest) GetRefundFee() int64 {
    return r._refundFee
}
// RefundReason Setter
// 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
func (r *TaobaoOpenmallRefundCreateRequest) SetRefundReason(_refundReason string) error {
    r._refundReason = _refundReason
    r.Set("refund_reason", _refundReason)
    return nil
}

// RefundReason Getter
func (r TaobaoOpenmallRefundCreateRequest) GetRefundReason() string {
    return r._refundReason
}
// RefundType Setter
// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
func (r *TaobaoOpenmallRefundCreateRequest) SetRefundType(_refundType string) error {
    r._refundType = _refundType
    r.Set("refund_type", _refundType)
    return nil
}

// RefundType Getter
func (r TaobaoOpenmallRefundCreateRequest) GetRefundType() string {
    return r._refundType
}
// Tid Setter
// 订单号
func (r *TaobaoOpenmallRefundCreateRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallRefundCreateRequest) GetTid() int64 {
    return r._tid
}
