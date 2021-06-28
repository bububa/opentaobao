package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量删除买赠商品 APIResponse
alibaba.wdk.marketing.buygift.item.remove.async

批量删除买赠商品
*/
type AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingBuygiftItemRemoveAsyncResponse `json:"alibaba_wdk_marketing_buygift_item_remove_async_response,omitempty"` 
    AlibabaWdkMarketingBuygiftItemRemoveAsyncResponse
}

/* model for simplify = false
type AlibabaWdkMarketingBuygiftItemRemoveAsyncResponse struct {

    // 结果信息
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingBuygiftItemRemoveAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
