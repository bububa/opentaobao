package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家电子凭证发码成功回调接口 APIResponse
taobao.vmarket.eticket.send

外部商家成功发码回调接口
*/
type TaobaoVmarketEticketSendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketSendResponse `json:"vmarket_eticket_send_response,omitempty"` 
    TaobaoVmarketEticketSendResponse
}

/* model for simplify = false
type TaobaoVmarketEticketSendResponse struct {

    // 0:失败；1:成功
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

}
*/

type TaobaoVmarketEticketSendResponse struct {

    // 0:失败；1:成功
    RetCode   int64 `json:"ret_code,omitempty"`

}
