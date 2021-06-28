package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
特价批量移除商品 APIResponse
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品
*/
type AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingDiscountItemRemoveAsyncResponse `json:"alibaba_wdk_marketing_discount_item_remove_async_response,omitempty"` 
    AlibabaWdkMarketingDiscountItemRemoveAsyncResponse
}

/* model for simplify = false
type AlibabaWdkMarketingDiscountItemRemoveAsyncResponse struct {

    // 结果信息
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingDiscountItemRemoveAsyncResponse struct {

    // 结果信息
    Result   *MarketResult `json:"result,omitempty"`

}
