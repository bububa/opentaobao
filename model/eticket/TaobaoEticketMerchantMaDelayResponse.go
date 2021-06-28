package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
凭证延期 APIResponse
taobao.eticket.merchant.ma.delay

订单延期
*/
type TaobaoEticketMerchantMaDelayAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoEticketMerchantMaDelayResponse `json:"eticket_merchant_ma_delay_response,omitempty"` 
    TaobaoEticketMerchantMaDelayResponse
}

/* model for simplify = false
type TaobaoEticketMerchantMaDelayResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    ResCode   int64 `json:"res_code,omitempty"`
    

    // 错误消息
    
    ResMsg   string `json:"res_msg,omitempty"`
    

}
*/

type TaobaoEticketMerchantMaDelayResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    ResCode   int64 `json:"res_code,omitempty"`

    // 错误消息
    ResMsg   string `json:"res_msg,omitempty"`

}
