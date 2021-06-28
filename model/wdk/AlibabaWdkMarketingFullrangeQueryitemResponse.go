package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
全场活动查询换购品 APIResponse
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品
*/
type AlibabaWdkMarketingFullrangeQueryitemAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingFullrangeQueryitemResponse `json:"alibaba_wdk_marketing_fullrange_queryitem_response,omitempty"` 
    AlibabaWdkMarketingFullrangeQueryitemResponse
}

/* model for simplify = false
type AlibabaWdkMarketingFullrangeQueryitemResponse struct {

    // 查询结果
    
    Result  *struct {
        MarketPageResult  *MarketPageResult `json:"market_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingFullrangeQueryitemResponse struct {

    // 查询结果
    Result   *MarketPageResult `json:"result,omitempty"`

}
