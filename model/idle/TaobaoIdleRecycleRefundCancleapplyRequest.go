package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收取消退款申请V2 API请求
taobao.idle.recycle.refund.cancleapply

回收商的回收订单取消退款申请
*/
type TaobaoIdleRecycleRefundCancleapplyRequest struct {
    model.Params
    // 订单号
    _bizOrderId   int64
}

// 初始化TaobaoIdleRecycleRefundCancleapplyRequest对象
func NewTaobaoIdleRecycleRefundCancleapplyRequest() *TaobaoIdleRecycleRefundCancleapplyRequest{
    return &TaobaoIdleRecycleRefundCancleapplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.cancleapply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *TaobaoIdleRecycleRefundCancleapplyRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoIdleRecycleRefundCancleapplyRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
