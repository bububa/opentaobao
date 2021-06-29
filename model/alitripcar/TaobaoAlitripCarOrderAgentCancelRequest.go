package alitripcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机或客服取消订单 APIRequest
taobao.alitrip.car.order.agent.cancel

司机或客服取消订单后通知飞猪订单取消
*/
type TaobaoAlitripCarOrderAgentCancelRequest struct {
    model.Params

    // 取消对象
    paramOrderCancel   *OrderCancel 

}

func NewTaobaoAlitripCarOrderAgentCancelRequest() *TaobaoAlitripCarOrderAgentCancelRequest{
    return &TaobaoAlitripCarOrderAgentCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarOrderAgentCancelRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.agent.cancel"
}

func (r TaobaoAlitripCarOrderAgentCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarOrderAgentCancelRequest) SetParamOrderCancel(paramOrderCancel *OrderCancel) error {
    r.paramOrderCancel = paramOrderCancel
    r.Set("param_order_cancel", paramOrderCancel)
    return nil
}

func (r TaobaoAlitripCarOrderAgentCancelRequest) GetParamOrderCancel() *OrderCancel {
    return r.paramOrderCancel
}

