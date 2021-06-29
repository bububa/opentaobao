package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退货 API请求
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。
*/
type TaobaoRpReturngoodsRefuseRequest struct {
    model.Params
    // 退款编号
    refundId   int64
    // 退款服务状态，售后或者售中
    refundPhase   string
    // 退款版本号
    refundVersion   int64
    // 拒绝退货凭证图片，必须图片格式，大小不能超过5M
    refuseProof   []*model.File
    // 拒绝原因编号，会提供拒绝原因列表供选择
    refuseReasonId   int64
}

// 初始化TaobaoRpReturngoodsRefuseRequest对象
func NewTaobaoRpReturngoodsRefuseRequest() *TaobaoRpReturngoodsRefuseRequest{
    return &TaobaoRpReturngoodsRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefuseRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundId() int64 {
    return r.refundId
}
// RefundPhase Setter
// 退款服务状态，售后或者售中
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundPhase() string {
    return r.refundPhase
}
// RefundVersion Setter
// 退款版本号
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundVersion(refundVersion int64) error {
    r.refundVersion = refundVersion
    r.Set("refund_version", refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundVersion() int64 {
    return r.refundVersion
}
// RefuseProof Setter
// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseProof(refuseProof []*model.File) error {
    r.refuseProof = refuseProof
    r.Set("refuse_proof", refuseProof)
    return nil
}

// RefuseProof Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseProof() []*model.File {
    return r.refuseProof
}
// RefuseReasonId Setter
// 拒绝原因编号，会提供拒绝原因列表供选择
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseReasonId(refuseReasonId int64) error {
    r.refuseReasonId = refuseReasonId
    r.Set("refuse_reason_id", refuseReasonId)
    return nil
}

// RefuseReasonId Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseReasonId() int64 {
    return r.refuseReasonId
}
