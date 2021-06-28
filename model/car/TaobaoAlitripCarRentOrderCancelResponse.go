package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
租车-取消订单 APIResponse
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单.
*/
type TaobaoAlitripCarRentOrderCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripCarRentOrderCancelResponse `json:"alitrip_car_rent_order_cancel_response,omitempty"` 
    TaobaoAlitripCarRentOrderCancelResponse
}

/* model for simplify = false
type TaobaoAlitripCarRentOrderCancelResponse struct {

    // 错误消息
    
    Message   string `json:"message,omitempty"`
    

    // 扩展对象
    
    Models   string `json:"models,omitempty"`
    

    // 结果对象
    
    Model   string `json:"model,omitempty"`
    

    // 结果码
    
    C   int64 `json:"c,omitempty"`
    

}
*/

type TaobaoAlitripCarRentOrderCancelResponse struct {

    // 错误消息
    Message   string `json:"message,omitempty"`

    // 扩展对象
    Models   string `json:"models,omitempty"`

    // 结果对象
    Model   string `json:"model,omitempty"`

    // 结果码
    C   int64 `json:"c,omitempty"`

}
