package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广 按照id操作推广计划的产品，包括新增，删除和更新 APIResponse
alibaba.scbp.target.ad.plan.update.products

定向推广 按照id操作推广计划的产品，包括新增，删除和更新
*/
type AlibabaScbpTargetAdPlanUpdateProductsAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanUpdateProductsResponse `json:"alibaba_scbp_target_ad_plan_update_products_response,omitempty"` 
    AlibabaScbpTargetAdPlanUpdateProductsResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanUpdateProductsResponse struct {

    // 操作成功的商品ID列表
    
    ProductIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"product_id_list,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanUpdateProductsResponse struct {

    // 操作成功的商品ID列表
    ProductIdList   []int64 `json:"product_id_list,omitempty"`

}
