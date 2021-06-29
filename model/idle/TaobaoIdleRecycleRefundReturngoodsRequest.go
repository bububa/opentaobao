package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收退货V2 APIRequest
taobao.idle.recycle.refund.returngoods

回收商买家退货，填写退货运单号
*/
type TaobaoIdleRecycleRefundReturngoodsRequest struct {
    model.Params

    // 退货
    param0   *RecycleReturnGoodsRequest 

}

func NewTaobaoIdleRecycleRefundReturngoodsRequest() *TaobaoIdleRecycleRefundReturngoodsRequest{
    return &TaobaoIdleRecycleRefundReturngoodsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIdleRecycleRefundReturngoodsRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.returngoods"
}

func (r TaobaoIdleRecycleRefundReturngoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIdleRecycleRefundReturngoodsRequest) SetParam0(param0 *RecycleReturnGoodsRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoIdleRecycleRefundReturngoodsRequest) GetParam0() *RecycleReturnGoodsRequest {
    return r.param0
}

