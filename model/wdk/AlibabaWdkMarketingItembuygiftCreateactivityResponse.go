package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建买赠活动 APIResponse
alibaba.wdk.marketing.itembuygift.createactivity

创建买赠活动
*/
type AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItembuygiftCreateactivityResponse `json:"alibaba_wdk_marketing_itembuygift_createactivity_response,omitempty"` 
    AlibabaWdkMarketingItembuygiftCreateactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItembuygiftCreateactivityResponse struct {

    // 创建活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItembuygiftCreateactivityResponse struct {

    // 创建活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
