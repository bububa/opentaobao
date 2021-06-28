package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
配送拦截接口 APIResponse
taobao.qimen.order.callback

配送拦截
*/
type TaobaoQimenOrderCallbackAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderCallbackResponse `json:"qimen_order_callback_response,omitempty"` 
    TaobaoQimenOrderCallbackResponse
}

/* model for simplify = false
type TaobaoQimenOrderCallbackResponse struct {

    // 
    
    Response  *struct {
        OrderCallbackResponseDO  *OrderCallbackResponseDO `json:"order_callback_response_do,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenOrderCallbackResponse struct {

    // 
    Response   *OrderCallbackResponseDO `json:"response,omitempty"`

}
