package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建商品特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动
*/
type AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItemdiscountCreateactivityResponse `json:"alibaba_wdk_marketing_itemdiscount_createactivity_response,omitempty"`
}

type AlibabaWdkMarketingItemdiscountCreateactivityResponse struct {

    // 创建活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}