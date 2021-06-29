package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收交易退款申请V2 API请求
taobao.idle.recycle.refund.apply

回收商买家申请退款
*/
type TaobaoIdleRecycleRefundApplyRequest struct {
    model.Params
    // 退款申请
    _refundApply   *RecycleRefundTopRequest
}

// 初始化TaobaoIdleRecycleRefundApplyRequest对象
func NewTaobaoIdleRecycleRefundApplyRequest() *TaobaoIdleRecycleRefundApplyRequest{
    return &TaobaoIdleRecycleRefundApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundApplyRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundApply Setter
// 退款申请
func (r *TaobaoIdleRecycleRefundApplyRequest) SetRefundApply(_refundApply *RecycleRefundTopRequest) error {
    r._refundApply = _refundApply
    r.Set("refund_apply", _refundApply)
    return nil
}

// RefundApply Getter
func (r TaobaoIdleRecycleRefundApplyRequest) GetRefundApply() *RecycleRefundTopRequest {
    return r._refundApply
}
