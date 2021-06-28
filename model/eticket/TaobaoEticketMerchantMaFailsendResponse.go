package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码商发码失败回调接口 APIResponse
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoEticketMerchantMaFailsendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoEticketMerchantMaFailsendResponse `json:"eticket_merchant_ma_failsend_response,omitempty"` 
    TaobaoEticketMerchantMaFailsendResponse
}

/* model for simplify = false
type TaobaoEticketMerchantMaFailsendResponse struct {

    // 回复参数
    
    RespBody  *struct {
        SendFailCallbackResp  *SendFailCallbackResp `json:"send_fail_callback_resp,omitempty"`
    } `json:"resp_body,omitempty"`
    

    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty"`
    

    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty"`
    

}
*/

type TaobaoEticketMerchantMaFailsendResponse struct {

    // 回复参数
    RespBody   *SendFailCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
