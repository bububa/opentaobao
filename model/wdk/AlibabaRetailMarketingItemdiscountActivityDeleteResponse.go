package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单品特价活动【同城零售】 APIResponse
alibaba.retail.marketing.itemdiscount.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItemdiscountActivityDeleteResponse `json:"alibaba_retail_marketing_itemdiscount_activity_delete_response,omitempty"`
}

type AlibabaRetailMarketingItemdiscountActivityDeleteResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
