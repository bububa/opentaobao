package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
司机应答API APIResponse
taobao.alitrip.car.order.confirm

航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
*/
type TaobaoAlitripCarOrderConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripCarOrderConfirmResponse `json:"alitrip_car_order_confirm_response,omitempty"` 
    TaobaoAlitripCarOrderConfirmResponse
}

/* model for simplify = false
type TaobaoAlitripCarOrderConfirmResponse struct {

    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty"`
    

    // 其它数据
    
    Data   string `json:"data,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

type TaobaoAlitripCarOrderConfirmResponse struct {

    // 错误码
    MessageCode   int64 `json:"message_code,omitempty"`

    // 其它数据
    Data   string `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

}
