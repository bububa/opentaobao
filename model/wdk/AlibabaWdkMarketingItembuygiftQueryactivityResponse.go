package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询买赠活动 APIResponse
alibaba.wdk.marketing.itembuygift.queryactivity

查询买赠活动
*/
type AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItembuygiftQueryactivityResponse `json:"alibaba_wdk_marketing_itembuygift_queryactivity_response,omitempty"` 
    AlibabaWdkMarketingItembuygiftQueryactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItembuygiftQueryactivityResponse struct {

    // 查询返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItembuygiftQueryactivityResponse struct {

    // 查询返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
