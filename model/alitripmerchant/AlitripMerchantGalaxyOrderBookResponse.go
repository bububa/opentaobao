package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单预订接口 APIResponse
alitrip.merchant.galaxy.order.book

为星河酒店解决方案的C端用户提供酒店预订能力
*/
type AlitripMerchantGalaxyOrderBookAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyOrderBookResponse
}

type AlitripMerchantGalaxyOrderBookResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_book_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
