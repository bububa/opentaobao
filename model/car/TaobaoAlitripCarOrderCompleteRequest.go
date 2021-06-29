package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务完成API API请求
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息
*/
type TaobaoAlitripCarOrderCompleteRequest struct {
    model.Params
    // 服务完成API
    paramOrderComplete   *OrderComplete
}

// 初始化TaobaoAlitripCarOrderCompleteRequest对象
func NewTaobaoAlitripCarOrderCompleteRequest() *TaobaoAlitripCarOrderCompleteRequest{
    return &TaobaoAlitripCarOrderCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderCompleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.complete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderComplete Setter
// 服务完成API
func (r *TaobaoAlitripCarOrderCompleteRequest) SetParamOrderComplete(paramOrderComplete *OrderComplete) error {
    r.paramOrderComplete = paramOrderComplete
    r.Set("param_order_complete", paramOrderComplete)
    return nil
}

// ParamOrderComplete Getter
func (r TaobaoAlitripCarOrderCompleteRequest) GetParamOrderComplete() *OrderComplete {
    return r.paramOrderComplete
}
