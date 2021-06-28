package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品池活动【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新
*/
type AlibabaRetailMarketingItempoolActivityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailMarketingItempoolActivityUpdateResponse `json:"alibaba_retail_marketing_itempool_activity_update_response,omitempty"` 
    AlibabaRetailMarketingItempoolActivityUpdateResponse
}

/* model for simplify = false
type AlibabaRetailMarketingItempoolActivityUpdateResponse struct {

    // 操作结果
    
    Result  *struct {
        OctopusOpenResult  *OctopusOpenResult `json:"octopus_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailMarketingItempoolActivityUpdateResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
