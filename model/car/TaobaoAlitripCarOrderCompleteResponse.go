package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务完成API APIResponse
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息
*/
type TaobaoAlitripCarOrderCompleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripCarOrderCompleteResponse `json:"alitrip_car_order_complete_response,omitempty"` 
    TaobaoAlitripCarOrderCompleteResponse
}

/* model for simplify = false
type TaobaoAlitripCarOrderCompleteResponse struct {

    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty"`
    

    // 其它数据
    
    Data   string `json:"data,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

type TaobaoAlitripCarOrderCompleteResponse struct {

    // 错误码
    MessageCode   int64 `json:"message_code,omitempty"`

    // 其它数据
    Data   string `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

}
