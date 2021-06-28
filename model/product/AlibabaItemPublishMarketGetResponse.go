package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商家可发布商品的市场信息 APIResponse
alibaba.item.publish.market.get

获取商家可发布商品的市场信息
*/
type AlibabaItemPublishMarketGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaItemPublishMarketGetResponse `json:"alibaba_item_publish_market_get_response,omitempty"` 
    AlibabaItemPublishMarketGetResponse
}

/* model for simplify = false
type AlibabaItemPublishMarketGetResponse struct {

    // 商家可发布的市场列表,多个以逗号(,)分隔
    
    Markets   string `json:"markets,omitempty"`
    

}
*/

type AlibabaItemPublishMarketGetResponse struct {

    // 商家可发布的市场列表,多个以逗号(,)分隔
    Markets   string `json:"markets,omitempty"`

}
