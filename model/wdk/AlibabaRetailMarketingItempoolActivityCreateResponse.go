package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.create

同城零售商品池活动创建
*/
type AlibabaRetailMarketingItempoolActivityCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailMarketingItempoolActivityCreateResponse `json:"alibaba_retail_marketing_itempool_activity_create_response,omitempty"` 
    AlibabaRetailMarketingItempoolActivityCreateResponse
}

/* model for simplify = false
type AlibabaRetailMarketingItempoolActivityCreateResponse struct {

    // 操作结果
    
    Result  *struct {
        OctopusOpenResult  *OctopusOpenResult `json:"octopus_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailMarketingItempoolActivityCreateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
