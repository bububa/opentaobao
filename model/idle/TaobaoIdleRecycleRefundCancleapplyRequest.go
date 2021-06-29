package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收取消退款申请V2 APIRequest
taobao.idle.recycle.refund.cancleapply

回收商的回收订单取消退款申请
*/
type TaobaoIdleRecycleRefundCancleapplyRequest struct {
    model.Params

    // 订单号
    bizOrderId   int64 

}

func NewTaobaoIdleRecycleRefundCancleapplyRequest() *TaobaoIdleRecycleRefundCancleapplyRequest{
    return &TaobaoIdleRecycleRefundCancleapplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.cancleapply"
}

func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIdleRecycleRefundCancleapplyRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

