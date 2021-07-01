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
type TaobaoRpReturngoodsRefuseAPIRequest struct {
    model.Params
    // 退款编号
    _refundId   int64
    // 退款服务状态，售后或者售中
    _refundPhase   string
    // 退款版本号
    _refundVersion   int64
    // 拒绝退货凭证图片，必须图片格式，大小不能超过5M
    _refuseProof   *model.File
    // 拒绝原因编号，会提供拒绝原因列表供选择
    _refuseReasonId   int64
}

// 初始化TaobaoRpReturngoodsRefuseAPIRequest对象
func NewTaobaoRpReturngoodsRefuseRequest() *TaobaoRpReturngoodsRefuseAPIRequest{
    return &TaobaoRpReturngoodsRefuseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundId() int64 {
    return r._refundId
}
// RefundPhase Setter
// 退款服务状态，售后或者售中
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundPhase(_refundPhase string) error {
    r._refundPhase = _refundPhase
    r.Set("refund_phase", _refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundPhase() string {
    return r._refundPhase
}
// RefundVersion Setter
// 退款版本号
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefundVersion(_refundVersion int64) error {
    r._refundVersion = _refundVersion
    r.Set("refund_version", _refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefundVersion() int64 {
    return r._refundVersion
}
// RefuseProof Setter
// 拒绝退货凭证图片，必须图片格式，大小不能超过5M
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefuseProof(_refuseProof *model.File) error {
    r._refuseProof = _refuseProof
    r.Set("refuse_proof", _refuseProof)
    return nil
}

// RefuseProof Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefuseProof() *model.File {
    return r._refuseProof
}
// RefuseReasonId Setter
// 拒绝原因编号，会提供拒绝原因列表供选择
func (r *TaobaoRpReturngoodsRefuseAPIRequest) SetRefuseReasonId(_refuseReasonId int64) error {
    r._refuseReasonId = _refuseReasonId
    r.Set("refuse_reason_id", _refuseReasonId)
    return nil
}

// RefuseReasonId Getter
func (r TaobaoRpReturngoodsRefuseAPIRequest) GetRefuseReasonId() int64 {
    return r._refuseReasonId
}
