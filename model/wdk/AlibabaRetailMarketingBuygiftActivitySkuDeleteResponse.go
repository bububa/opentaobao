package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动商品 APIResponse
alibaba.retail.marketing.buygift.activity.sku.delete

删除单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse `json:"alibaba_retail_marketing_buygift_activity_sku_delete_response,omitempty"`
}

type AlibabaRetailMarketingBuygiftActivitySkuDeleteResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
