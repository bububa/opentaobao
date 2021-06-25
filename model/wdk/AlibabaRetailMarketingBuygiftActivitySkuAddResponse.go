package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加单品买赠活动商品 APIResponse
alibaba.retail.marketing.buygift.activity.sku.add

新增或更新单品买赠活动商品信息【同城零售】
*/
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaRetailMarketingBuygiftActivitySkuAddResponse `json:"alibaba_retail_marketing_buygift_activity_sku_add_response,omitempty"`
}

type AlibabaRetailMarketingBuygiftActivitySkuAddResponse struct {

    // 操作结果
    Result   *OctopusOpenResult `json:"result,omitempty"`

}
