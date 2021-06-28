package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证重发回调接口 APIResponse
taobao.eticket.merchant.ma.resend

码商重发电子凭证回调接口
*/
type TaobaoEticketMerchantMaResendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoEticketMerchantMaResendResponse `json:"eticket_merchant_ma_resend_response,omitempty"` 
    TaobaoEticketMerchantMaResendResponse
}

/* model for simplify = false
type TaobaoEticketMerchantMaResendResponse struct {

    // 回复参数
    
    RespBody  *struct {
        SendMaCallbackResp  *SendMaCallbackResp `json:"send_ma_callback_resp,omitempty"`
    } `json:"resp_body,omitempty"`
    

    // 子结果码
    
    RetCode   string `json:"ret_code,omitempty"`
    

    // 子结果信息
    
    RetMsg   string `json:"ret_msg,omitempty"`
    

}
*/

type TaobaoEticketMerchantMaResendResponse struct {

    // 回复参数
    RespBody   *SendMaCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
