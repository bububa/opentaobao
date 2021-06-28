package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过店铺id取得短链 APIResponse
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链
*/
type TaobaoWirelessBuntingShopShorturlCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWirelessBuntingShopShorturlCreateResponse `json:"wireless_bunting_shop_shorturl_create_response,omitempty"` 
    TaobaoWirelessBuntingShopShorturlCreateResponse
}

/* model for simplify = false
type TaobaoWirelessBuntingShopShorturlCreateResponse struct {

    // 短链
    
    Shorturl   string `json:"shorturl,omitempty"`
    

}
*/

type TaobaoWirelessBuntingShopShorturlCreateResponse struct {

    // 短链
    Shorturl   string `json:"shorturl,omitempty"`

}
