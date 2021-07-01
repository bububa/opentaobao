package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票下单接口 API请求
taobao.bus.order.set

提供给汽车票商家进行下单
*/
type TaobaoBusOrderSetAPIRequest struct {
    model.Params
    // 下单参数
    _paramB2BCreateOrderRQ   *B2BCreateOrderRq
}

// 初始化TaobaoBusOrderSetAPIRequest对象
func NewTaobaoBusOrderSetRequest() *TaobaoBusOrderSetAPIRequest{
    return &TaobaoBusOrderSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusOrderSetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.order.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusOrderSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamB2BCreateOrderRQ Setter
// 下单参数
func (r *TaobaoBusOrderSetAPIRequest) SetParamB2BCreateOrderRQ(_paramB2BCreateOrderRQ *B2BCreateOrderRq) error {
    r._paramB2BCreateOrderRQ = _paramB2BCreateOrderRQ
    r.Set("param_b2_b_create_order_r_q", _paramB2BCreateOrderRQ)
    return nil
}

// ParamB2BCreateOrderRQ Getter
func (r TaobaoBusOrderSetAPIRequest) GetParamB2BCreateOrderRQ() *B2BCreateOrderRq {
    return r._paramB2BCreateOrderRQ
}
