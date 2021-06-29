package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机或客服取消订单 API请求
taobao.alitrip.car.order.agent.cancel

司机或客服取消订单后通知飞猪订单取消
*/
type TaobaoAlitripCarOrderAgentCancelRequest struct {
    model.Params
    // 取消对象
    paramOrderCancel   *OrderCancel
}

// 初始化TaobaoAlitripCarOrderAgentCancelRequest对象
func NewTaobaoAlitripCarOrderAgentCancelRequest() *TaobaoAlitripCarOrderAgentCancelRequest{
    return &TaobaoAlitripCarOrderAgentCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderAgentCancelRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.agent.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderAgentCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderCancel Setter
// 取消对象
func (r *TaobaoAlitripCarOrderAgentCancelRequest) SetParamOrderCancel(paramOrderCancel *OrderCancel) error {
    r.paramOrderCancel = paramOrderCancel
    r.Set("param_order_cancel", paramOrderCancel)
    return nil
}

// ParamOrderCancel Getter
func (r TaobaoAlitripCarOrderAgentCancelRequest) GetParamOrderCancel() *OrderCancel {
    return r.paramOrderCancel
}
