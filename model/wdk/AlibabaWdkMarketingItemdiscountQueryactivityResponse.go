package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查找特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
type AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItemdiscountQueryactivityResponse `json:"alibaba_wdk_marketing_itemdiscount_queryactivity_response,omitempty"`
}

type AlibabaWdkMarketingItemdiscountQueryactivityResponse struct {

    // 查询特价活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
