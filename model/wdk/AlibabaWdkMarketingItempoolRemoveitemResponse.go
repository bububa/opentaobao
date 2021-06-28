package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
移除商品池里面的商品 APIResponse
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品
*/
type AlibabaWdkMarketingItempoolRemoveitemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItempoolRemoveitemResponse `json:"alibaba_wdk_marketing_itempool_removeitem_response,omitempty"` 
    AlibabaWdkMarketingItempoolRemoveitemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItempoolRemoveitemResponse struct {

    // 移除商品返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItempoolRemoveitemResponse struct {

    // 移除商品返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
