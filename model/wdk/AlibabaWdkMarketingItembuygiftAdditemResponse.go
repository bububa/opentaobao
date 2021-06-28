package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加买赠活动商品。【注意，此接口暂不支持并发！】 APIResponse
alibaba.wdk.marketing.itembuygift.additem

增加买赠活动商品。【注意，此接口暂不支持并发！】
*/
type AlibabaWdkMarketingItembuygiftAdditemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItembuygiftAdditemResponse `json:"alibaba_wdk_marketing_itembuygift_additem_response,omitempty"` 
    AlibabaWdkMarketingItembuygiftAdditemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItembuygiftAdditemResponse struct {

    // 商品报名活动的返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItembuygiftAdditemResponse struct {

    // 商品报名活动的返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
