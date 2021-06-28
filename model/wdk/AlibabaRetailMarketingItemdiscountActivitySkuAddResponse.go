package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
特价活动新增商品 APIResponse
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaRetailMarketingItemdiscountActivitySkuAddResponse `json:"alibaba_retail_marketing_itemdiscount_activity_sku_add_response,omitempty"` 
    AlibabaRetailMarketingItemdiscountActivitySkuAddResponse
}

/* model for simplify = false
type AlibabaRetailMarketingItemdiscountActivitySkuAddResponse struct {

    // 操作结果
    
    Result  *struct {
        OctopusOpenResult  *OctopusOpenResult `json:"octopus_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaRetailMarketingItemdiscountActivitySkuAddResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
