package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除商品池活动商品【同城零售】 APIResponse
alibaba.retail.marketing.itempool.activity.sku.delete

删除商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItempoolActivitySkuDeleteResponse `json:"alibaba_retail_marketing_itempool_activity_sku_delete_response,omitempty"`
}

type AlibabaRetailMarketingItempoolActivitySkuDeleteResponse struct {

    // 出参
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
