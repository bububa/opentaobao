package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除买赠活动 APIResponse
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动
*/
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItembuygiftDeleteactivityResponse `json:"alibaba_wdk_marketing_itembuygift_deleteactivity_response,omitempty"` 
    AlibabaWdkMarketingItembuygiftDeleteactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItembuygiftDeleteactivityResponse struct {

    // 删除活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItembuygiftDeleteactivityResponse struct {

    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
