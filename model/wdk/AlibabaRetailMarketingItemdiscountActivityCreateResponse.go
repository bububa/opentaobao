package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建单品特价活动【同城零售】 APIResponse
alibaba.retail.marketing.itemdiscount.activity.create

同城零售单品特价活动创建
*/
type AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItemdiscountActivityCreateResponse `json:"alibaba_retail_marketing_itemdiscount_activity_create_response,omitempty"`
}

type AlibabaRetailMarketingItemdiscountActivityCreateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
