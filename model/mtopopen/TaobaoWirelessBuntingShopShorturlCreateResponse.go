package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过店铺id取得短链 APIResponse
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链
*/
type TaobaoWirelessBuntingShopShorturlCreateAPIResponse struct {
    model.CommonResponse
    TaobaoWirelessBuntingShopShorturlCreateResponse
}

type TaobaoWirelessBuntingShopShorturlCreateResponse struct {
    XMLName xml.Name `xml:"wireless_bunting_shop_shorturl_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 短链
    
    Shorturl   string `json:"shorturl,omitempty" xml:"shorturl,omitempty"`

    
}
