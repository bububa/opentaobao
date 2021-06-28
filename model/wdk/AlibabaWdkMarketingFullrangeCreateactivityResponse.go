package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建全场活动 APIResponse
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动
*/
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingFullrangeCreateactivityResponse `json:"alibaba_wdk_marketing_fullrange_createactivity_response,omitempty"` 
    AlibabaWdkMarketingFullrangeCreateactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingFullrangeCreateactivityResponse struct {

    // 创建活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingFullrangeCreateactivityResponse struct {

    // 创建活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
