package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收退款详情V2 APIRequest
taobao.idle.recycle.refund.detail

回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
*/
type TaobaoIdleRecycleRefundDetailRequest struct {
    model.Params

    // 订单号
    bizOrderId   int64 

}

func NewTaobaoIdleRecycleRefundDetailRequest() *TaobaoIdleRecycleRefundDetailRequest{
    return &TaobaoIdleRecycleRefundDetailRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIdleRecycleRefundDetailRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.detail"
}

func (r TaobaoIdleRecycleRefundDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIdleRecycleRefundDetailRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoIdleRecycleRefundDetailRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

