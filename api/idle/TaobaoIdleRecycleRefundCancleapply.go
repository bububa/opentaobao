package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
闲鱼回收取消退款申请V2 
taobao.idle.recycle.refund.cancleapply

回收商的回收订单取消退款申请
*/
func TaobaoIdleRecycleRefundCancleapply(clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundCancleapplyAPIRequest, session string) (*idle.TaobaoIdleRecycleRefundCancleapplyAPIResponse, error) {
    var resp idle.TaobaoIdleRecycleRefundCancleapplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
