package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的核销接口 APIResponse
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
*/
type TaobaoVmarketEticketAuthConsumeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVmarketEticketAuthConsumeResponse `json:"vmarket_eticket_auth_consume_response,omitempty"` 
    TaobaoVmarketEticketAuthConsumeResponse
}

/* model for simplify = false
type TaobaoVmarketEticketAuthConsumeResponse struct {

    // 1:可以进行核销码操作
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 商品标题
    
    ItemTitle   string `json:"item_title,omitempty"`
    

    // 订单ID
    
    OrderId   int64 `json:"order_id,omitempty"`
    

    // 淘宝卖家ID
    
    TaobaoSid   int64 `json:"taobao_sid,omitempty"`
    

    // 淘宝卖家旺旺名称
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

}
*/

type TaobaoVmarketEticketAuthConsumeResponse struct {

    // 1:可以进行核销码操作
    RetCode   int64 `json:"ret_code,omitempty"`

    // 商品标题
    ItemTitle   string `json:"item_title,omitempty"`

    // 订单ID
    OrderId   int64 `json:"order_id,omitempty"`

    // 淘宝卖家ID
    TaobaoSid   int64 `json:"taobao_sid,omitempty"`

    // 淘宝卖家旺旺名称
    SellerNick   string `json:"seller_nick,omitempty"`

}
