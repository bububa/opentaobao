package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加商品池活动 APIResponse
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动
*/
type AlibabaWdkMarketingItempoolCreateactivityAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingItempoolCreateactivityResponse `json:"alibaba_wdk_marketing_itempool_createactivity_response,omitempty"` 
    AlibabaWdkMarketingItempoolCreateactivityResponse
}

/* model for simplify = false
type AlibabaWdkMarketingItempoolCreateactivityResponse struct {

    // 创建活动返回结果
    
    Result  *struct {
        MarketResult  *MarketResult `json:"market_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingItempoolCreateactivityResponse struct {

    // 创建活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
