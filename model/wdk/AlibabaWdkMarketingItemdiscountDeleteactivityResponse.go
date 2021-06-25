package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除商品特价活动 APIResponse
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动
*/
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingItemdiscountDeleteactivityResponse `json:"alibaba_wdk_marketing_itemdiscount_deleteactivity_response,omitempty"`
}

type AlibabaWdkMarketingItemdiscountDeleteactivityResponse struct {

    // 删除活动返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
