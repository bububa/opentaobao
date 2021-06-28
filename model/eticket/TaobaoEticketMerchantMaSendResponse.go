package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码商发码成功回调接口 APIResponse
taobao.eticket.merchant.ma.send

码商发码成功回调接口
*/
type TaobaoEticketMerchantMaSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoEticketMerchantMaSendResponse `json:"eticket_merchant_ma_send_response,omitempty"` 
    TaobaoEticketMerchantMaSendResponse
}

/* model for simplify = false
type TaobaoEticketMerchantMaSendResponse struct {

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

type TaobaoEticketMerchantMaSendResponse struct {

    // 回复参数
    RespBody   *SendMaCallbackResp `json:"resp_body,omitempty"`

    // 子结果码
    RetCode   string `json:"ret_code,omitempty"`

    // 子结果信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
