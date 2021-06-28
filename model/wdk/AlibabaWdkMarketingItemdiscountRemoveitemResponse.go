package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
移除报名的商品 APIResponse
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除
*/
type AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItemdiscountRemoveitemResponse `json:"alibaba_wdk_marketing_itemdiscount_removeitem_response,omitempty"` 
    AlibabaWdkMarketingItemdiscountRemoveitemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItemdiscountRemoveitemResponse struct {

    // 移除商品返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItemdiscountRemoveitemResponse struct {

    // 移除商品返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
