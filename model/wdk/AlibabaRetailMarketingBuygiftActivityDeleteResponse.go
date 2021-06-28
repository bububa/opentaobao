package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailMarketingBuygiftActivityDeleteResponse `json:"alibaba_retail_marketing_buygift_activity_delete_response,omitempty"` 
    AlibabaRetailMarketingBuygiftActivityDeleteResponse
}

/* model for simplify = false
type AlibabaRetailMarketingBuygiftActivityDeleteResponse struct {

    // 操作结果
    
    Result  *struct {
        OctopusOpenResult  *OctopusOpenResult `json:"octopus_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailMarketingBuygiftActivityDeleteResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
