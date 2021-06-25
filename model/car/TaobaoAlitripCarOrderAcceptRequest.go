package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认接单 APIRequest
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息
*/
type TaobaoAlitripCarOrderAcceptRequest struct {
    model.Params

    // 确认订单请求
    paramOrderAccept   *OrderAccept 

}

func NewTaobaoAlitripCarOrderAcceptRequest() *TaobaoAlitripCarOrderAcceptRequest{
    return &TaobaoAlitripCarOrderAcceptRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarOrderAcceptRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.accept"
}

func (r TaobaoAlitripCarOrderAcceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarOrderAcceptRequest) SetParamOrderAccept(paramOrderAccept *OrderAccept) error {
    r.paramOrderAccept = paramOrderAccept
    r.Set("param_order_accept", paramOrderAccept)
    return nil
}

func (r TaobaoAlitripCarOrderAcceptRequest) GetParamOrderAccept() *OrderAccept {
    return r.paramOrderAccept
}

