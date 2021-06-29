package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机应答API API请求
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
*/
type TaobaoAlitripCarOrderConfirmRequest struct {
    model.Params
    // 请求对象
    paramOrderConfirm   *OrderConfirm
}

// 初始化TaobaoAlitripCarOrderConfirmRequest对象
func NewTaobaoAlitripCarOrderConfirmRequest() *TaobaoAlitripCarOrderConfirmRequest{
    return &TaobaoAlitripCarOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderConfirm Setter
// 请求对象
func (r *TaobaoAlitripCarOrderConfirmRequest) SetParamOrderConfirm(paramOrderConfirm *OrderConfirm) error {
    r.paramOrderConfirm = paramOrderConfirm
    r.Set("param_order_confirm", paramOrderConfirm)
    return nil
}

// ParamOrderConfirm Getter
func (r TaobaoAlitripCarOrderConfirmRequest) GetParamOrderConfirm() *OrderConfirm {
    return r.paramOrderConfirm
}
