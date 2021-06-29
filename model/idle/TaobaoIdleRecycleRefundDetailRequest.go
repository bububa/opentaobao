package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收退款详情V2 API请求
taobao.idle.recycle.refund.detail

回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
*/
type TaobaoIdleRecycleRefundDetailRequest struct {
    model.Params
    // 订单号
    bizOrderId   int64
}

// 初始化TaobaoIdleRecycleRefundDetailRequest对象
func NewTaobaoIdleRecycleRefundDetailRequest() *TaobaoIdleRecycleRefundDetailRequest{
    return &TaobaoIdleRecycleRefundDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundDetailRequest) GetApiMethodName() string {
    return "taobao.idle.recycle.refund.detail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单号
func (r *TaobaoIdleRecycleRefundDetailRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TaobaoIdleRecycleRefundDetailRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}
