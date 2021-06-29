package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除档期 APIResponse
alibaba.price.promotion.item.delete

盒马帮批量删除档期商品
*/
type AlibabaPricePromotionItemDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaPricePromotionItemDeleteResponse
}

type AlibabaPricePromotionItemDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_price_promotion_item_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaPricePromotionItemDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
