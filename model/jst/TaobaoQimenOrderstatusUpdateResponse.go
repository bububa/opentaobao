package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订单状态更新接口 APIResponse
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
*/
type TaobaoQimenOrderstatusUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenOrderstatusUpdateResponse `json:"qimen_orderstatus_update_response,omitempty"` 
    TaobaoQimenOrderstatusUpdateResponse
}

/* model for simplify = false
type TaobaoQimenOrderstatusUpdateResponse struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoQimenOrderstatusUpdateResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
