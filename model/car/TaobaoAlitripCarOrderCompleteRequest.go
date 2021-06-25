package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务完成API APIRequest
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息
*/
type TaobaoAlitripCarOrderCompleteRequest struct {
    model.Params

    // 服务完成API
    paramOrderComplete   *OrderComplete 

}

func NewTaobaoAlitripCarOrderCompleteRequest() *TaobaoAlitripCarOrderCompleteRequest{
    return &TaobaoAlitripCarOrderCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarOrderCompleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.complete"
}

func (r TaobaoAlitripCarOrderCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarOrderCompleteRequest) SetParamOrderComplete(paramOrderComplete *OrderComplete) error {
    r.paramOrderComplete = paramOrderComplete
    r.Set("param_order_complete", paramOrderComplete)
    return nil
}

func (r TaobaoAlitripCarOrderCompleteRequest) GetParamOrderComplete() *OrderComplete {
    return r.paramOrderComplete
}

