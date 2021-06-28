package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取推广计划产品列表 APIResponse
alibaba.scbp.target.ad.plan.product.list.get

定向推广-获取推广计划产品列表
*/
type AlibabaScbpTargetAdPlanProductListGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanProductListGetResponse `json:"alibaba_scbp_target_ad_plan_product_list_get_response,omitempty"` 
    AlibabaScbpTargetAdPlanProductListGetResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanProductListGetResponse struct {

    // TopP4pQuickCampaignProductView
    
    ProductList  struct {
        TopP4pQuickCampaignProductView  []TopP4pQuickCampaignProductView `json:"top_p4p_quick_campaign_product_view,omitempty"`
    } `json:"product_list,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanProductListGetResponse struct {

    // TopP4pQuickCampaignProductView
    ProductList   []TopP4pQuickCampaignProductView `json:"product_list,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

}
