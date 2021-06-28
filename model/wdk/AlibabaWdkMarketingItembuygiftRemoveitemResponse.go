package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
移除买赠活动下的商品。【注意，此接口暂不支持并发！】 APIResponse
alibaba.wdk.marketing.itembuygift.removeitem

移除买赠活动下的商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItembuygiftRemoveitemResponse `json:"alibaba_wdk_marketing_itembuygift_removeitem_response,omitempty"` 
    AlibabaWdkMarketingItembuygiftRemoveitemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItembuygiftRemoveitemResponse struct {

    // 移除商品返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItembuygiftRemoveitemResponse struct {

    // 移除商品返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
