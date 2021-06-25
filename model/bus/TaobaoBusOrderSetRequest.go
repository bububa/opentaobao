package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票下单接口 APIRequest
taobao.bus.order.set

提供给汽车票商家进行下单
*/
type TaobaoBusOrderSetRequest struct {
    model.Params

    // 下单参数
    paramB2BCreateOrderRQ   *B2BCreateOrderRq 

}

func NewTaobaoBusOrderSetRequest() *TaobaoBusOrderSetRequest{
    return &TaobaoBusOrderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusOrderSetRequest) GetApiMethodName() string {
    return "taobao.bus.order.set"
}

func (r TaobaoBusOrderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusOrderSetRequest) SetParamB2BCreateOrderRQ(paramB2BCreateOrderRQ *B2BCreateOrderRq) error {
    r.paramB2BCreateOrderRQ = paramB2BCreateOrderRQ
    r.Set("param_b2_b_create_order_r_q", paramB2BCreateOrderRQ)
    return nil
}

func (r TaobaoBusOrderSetRequest) GetParamB2BCreateOrderRQ() *B2BCreateOrderRq {
    return r.paramB2BCreateOrderRQ
}

