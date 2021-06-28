package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无法发码回调 APIResponse
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoVmarketEticketFailsendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketFailsendResponse `json:"vmarket_eticket_failsend_response,omitempty"` 
    TaobaoVmarketEticketFailsendResponse
}

/* model for simplify = false
type TaobaoVmarketEticketFailsendResponse struct {

    // 成功
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

}
*/

type TaobaoVmarketEticketFailsendResponse struct {

    // 成功
    RetCode   int64 `json:"ret_code,omitempty"`

}
