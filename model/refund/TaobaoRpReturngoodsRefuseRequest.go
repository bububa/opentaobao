package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝退货 APIRequest
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

func NewTaobaoRpReturngoodsRefuseRequest() *TaobaoRpReturngoodsRefuseRequest{
    return &TaobaoRpReturngoodsRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRpReturngoodsRefuseRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.refuse"
}

func (r TaobaoRpReturngoodsRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoRpReturngoodsRefuseRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

func (r TaobaoRpReturngoodsRefuseRequest) GetRefundPhase() string {
    return r.refundPhase
}

func (r *TaobaoRpReturngoodsRefuseRequest) SetRefundVersion(refundVersion int64) error {
    r.refundVersion = refundVersion
    r.Set("refund_version", refundVersion)
    return nil
}

func (r TaobaoRpReturngoodsRefuseRequest) GetRefundVersion() int64 {
    return r.refundVersion
}

func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseProof(refuseProof []*model.File) error {
    r.refuseProof = refuseProof
    r.Set("refuse_proof", refuseProof)
    return nil
}

func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseProof() []*model.File {
    return r.refuseProof
}

func (r *TaobaoRpReturngoodsRefuseRequest) SetRefuseReasonId(refuseReasonId int64) error {
    r.refuseReasonId = refuseReasonId
    r.Set("refuse_reason_id", refuseReasonId)
    return nil
}

func (r TaobaoRpReturngoodsRefuseRequest) GetRefuseReasonId() int64 {
    return r.refuseReasonId
}

