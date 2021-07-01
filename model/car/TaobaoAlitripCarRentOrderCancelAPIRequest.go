package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
租车-取消订单 API请求
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单.
*/
type TaobaoAlitripCarRentOrderCancelAPIRequest struct {
    model.Params
    // 取消请求对象
    _param0   *RentProviderCancelRequest
}

// 初始化TaobaoAlitripCarRentOrderCancelAPIRequest对象
func NewTaobaoAlitripCarRentOrderCancelRequest() *TaobaoAlitripCarRentOrderCancelAPIRequest{
    return &TaobaoAlitripCarRentOrderCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.rent.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 取消请求对象
func (r *TaobaoAlitripCarRentOrderCancelAPIRequest) SetParam0(_param0 *RentProviderCancelRequest) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetParam0() *RentProviderCancelRequest {
    return r._param0
}
