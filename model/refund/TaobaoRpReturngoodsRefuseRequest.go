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
    _refundId   int64
    // 退款服务状态，售后或者售中
    _refundPhase   string
    // 退款版本号
    _refundVersion   int64
    // 拒绝退货凭证图片，必须图片格式，大小不能超过5M
    _refuseProof   []*model.File
    // 拒绝原因编号，会提供拒绝原因列表供选择
    _refuseReasonId   int64
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
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundId() int64 {
    return r._refundId
}
// RefundPhase Setter
// 退款服务状态，售后或者售中
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundPhase(_refundPhase string) error {
    r._refundPhase = _refundPhase
    r.Set("refund_phase", _refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundPhase() string {
    return r._refundPhase
}
// RefundVersion Setter
// 退款版本号
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundVersion(_refundVersion int64) error {
    r._refundVersion = _refundVersion
    r.Set("refund_version", _refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefundVersion() int64 {
    return r._refundVersion
}
// RefuseProof Setter
// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseProof(_refuseProof []*model.File) error {
    r._refuseProof = _refuseProof
    r.Set("refuse_proof", _refuseProof)
    return nil
}

// RefuseProof Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseProof() []*model.File {
    return r._refuseProof
}
// RefuseReasonId Setter
// 拒绝原因编号，会提供拒绝原因列表供选择
func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseReasonId(_refuseReasonId int64) error {
    r._refuseReasonId = _refuseReasonId
    r.Set("refuse_reason_id", _refuseReasonId)
    return nil
}

// RefuseReasonId Getter
func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseReasonId() int64 {
    return r._refuseReasonId
}
