package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票订单查询 APIRequest
taobao.bus.order.get

商家汽车票订单查询
*/
type TaobaoBusOrderGetRequest struct {
    model.Params

    // 订单查询对象
    paramB2BOrderQueryRQ   *B2BOrderQueryRq 

}

func NewTaobaoBusOrderGetRequest() *TaobaoBusOrderGetRequest{
    return &TaobaoBusOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusOrderGetRequest) GetApiMethodName() string {
    return "taobao.bus.order.get"
}

func (r TaobaoBusOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusOrderGetRequest) SetParamB2BOrderQueryRQ(paramB2BOrderQueryRQ *B2BOrderQueryRq) error {
    r.paramB2BOrderQueryRQ = paramB2BOrderQueryRQ
    r.Set("param_b2_b_order_query_r_q", paramB2BOrderQueryRQ)
    return nil
}

func (r TaobaoBusOrderGetRequest) GetParamB2BOrderQueryRQ() *B2BOrderQueryRq {
    return r.paramB2BOrderQueryRQ
}

