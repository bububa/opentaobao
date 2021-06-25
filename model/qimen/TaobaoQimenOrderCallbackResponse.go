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
    Response *TaobaoQimenOrderCallbackResponse `json:"taobao_qimen_order_callback_response,omitempty"`
}

type TaobaoQimenOrderCallbackResponse struct {

    // 
    Response   *OrderCallbackResponseDO `json:"response,omitempty"`

}
