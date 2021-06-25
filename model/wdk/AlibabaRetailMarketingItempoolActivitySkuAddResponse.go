package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品池活动新增商品 APIResponse
alibaba.retail.marketing.itempool.activity.sku.add

新增或更新商品池活动商品信息【同城零售】
*/
type AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingItempoolActivitySkuAddResponse `json:"alibaba_retail_marketing_itempool_activity_sku_add_response,omitempty"`
}

type AlibabaRetailMarketingItempoolActivitySkuAddResponse struct {

    // 出参
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
