package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新单品特价活动【同城零售】 APIResponse
alibaba.retail.marketing.itemdiscount.activity.update

同城零售单品特价活动更新
*/
type AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItemdiscountActivityUpdateResponse `json:"alibaba_retail_marketing_itemdiscount_activity_update_response,omitempty"`
}

type AlibabaRetailMarketingItemdiscountActivityUpdateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
