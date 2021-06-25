package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
租车-取消订单 APIRequest
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单.
*/
type TaobaoAlitripCarRentOrderCancelRequest struct {
    model.Params

    // 取消请求对象
    param0   *RentProviderCancelRequest 

}

func NewTaobaoAlitripCarRentOrderCancelRequest() *TaobaoAlitripCarRentOrderCancelRequest{
    return &TaobaoAlitripCarRentOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarRentOrderCancelRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.rent.order.cancel"
}

func (r TaobaoAlitripCarRentOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarRentOrderCancelRequest) SetParam0(param0 *RentProviderCancelRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAlitripCarRentOrderCancelRequest) GetParam0() *RentProviderCancelRequest {
    return r.param0
}

