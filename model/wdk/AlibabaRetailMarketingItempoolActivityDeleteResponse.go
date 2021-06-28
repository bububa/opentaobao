package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.delete

同城零售商品池活动删除
*/
type AlibabaRetailMarketingItempoolActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailMarketingItempoolActivityDeleteResponse `json:"alibaba_retail_marketing_itempool_activity_delete_response,omitempty"` 
    AlibabaRetailMarketingItempoolActivityDeleteResponse
}

/* model for simplify = false
type AlibabaRetailMarketingItempoolActivityDeleteResponse struct {

    // 出参
    
    Result  *struct {
        OctopusOpenResult  *OctopusOpenResult `json:"octopus_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailMarketingItempoolActivityDeleteResponse struct {

    // 出参
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
