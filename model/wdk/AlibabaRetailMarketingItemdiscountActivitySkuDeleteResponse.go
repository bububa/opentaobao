package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除特价活动商品 APIResponse
alibaba.retail.marketing.itemdiscount.activity.sku.delete

删除活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItemdiscountActivitySkuDeleteResponse `json:"alibaba_retail_marketing_itemdiscount_activity_sku_delete_response,omitempty"`
}

type AlibabaRetailMarketingItemdiscountActivitySkuDeleteResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
