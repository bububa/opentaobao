package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收交易退款申请V2 APIRequest
taobao.idle.recycle.refund.apply

回收商买家申请退款
*/
type TaobaoIdleRecycleRefundApplyRequest struct {
    model.Params

    // 退款申请
    refundApply   *RecycleRefundTopRequest 

}

func NewTaobaoIdleRecycleRefundApplyRequest() *TaobaoIdleRecycleRefundApplyRequest{
    return &TaobaoIdleRecycleRefundApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIdleRecycleRefundApplyRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.apply"
}

func (r TaobaoIdleRecycleRefundApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIdleRecycleRefundApplyRequest) SetRefundApply(refundApply *RecycleRefundTopRequest) error {
    r.refundApply = refundApply
    r.Set("refund_apply", refundApply)
    return nil
}

func (r TaobaoIdleRecycleRefundApplyRequest) GetRefundApply() *RecycleRefundTopRequest {
    return r.refundApply
}

