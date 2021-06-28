package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
业务重新触发发码短信 APIResponse
taobao.vmarket.eticket.flow.resend

业务重新触发发码短信
*/
type TaobaoVmarketEticketFlowResendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketFlowResendResponse `json:"vmarket_eticket_flow_resend_response,omitempty"` 
    TaobaoVmarketEticketFlowResendResponse
}

/* model for simplify = false
type TaobaoVmarketEticketFlowResendResponse struct {

    // 1成功;0失败
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 错误提示信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type TaobaoVmarketEticketFlowResendResponse struct {

    // 1成功;0失败
    RetCode   int64 `json:"ret_code,omitempty"`

    // 错误提示信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
