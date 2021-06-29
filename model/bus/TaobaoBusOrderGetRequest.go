package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票订单查询 API请求
taobao.bus.order.get

商家汽车票订单查询
*/
type TaobaoBusOrderGetRequest struct {
    model.Params
    // 订单查询对象
    paramB2BOrderQueryRQ   *B2BOrderQueryRq
}

// 初始化TaobaoBusOrderGetRequest对象
func NewTaobaoBusOrderGetRequest() *TaobaoBusOrderGetRequest{
    return &TaobaoBusOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusOrderGetRequest) GetApiMethodName() string {
    return "taobao.bus.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamB2BOrderQueryRQ Setter
// 订单查询对象
func (r *TaobaoBusOrderGetRequest) SetParamB2BOrderQueryRQ(paramB2BOrderQueryRQ *B2BOrderQueryRq) error {
    r.paramB2BOrderQueryRQ = paramB2BOrderQueryRQ
    r.Set("param_b2_b_order_query_r_q", paramB2BOrderQueryRQ)
    return nil
}

// ParamB2BOrderQueryRQ Getter
func (r TaobaoBusOrderGetRequest) GetParamB2BOrderQueryRQ() *B2BOrderQueryRq {
    return r.paramB2BOrderQueryRQ
}
