package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品池删除商品 APIResponse
alibaba.wdk.marketing.itempool.item.remove.async

新模型下删除商品
*/
type AlibabaWdkMarketingItempoolItemRemoveAsyncAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItempoolItemRemoveAsyncResponse `json:"alibaba_wdk_marketing_itempool_item_remove_async_response,omitempty"` 
    AlibabaWdkMarketingItempoolItemRemoveAsyncResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItempoolItemRemoveAsyncResponse struct {

    // 结果信息
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItempoolItemRemoveAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
