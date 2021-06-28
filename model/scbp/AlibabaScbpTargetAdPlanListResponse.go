package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-查询定向推广计划列表并返回计划基础信息 APIResponse
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息
*/
type AlibabaScbpTargetAdPlanListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanListResponse `json:"alibaba_scbp_target_ad_plan_list_response,omitempty"` 
    AlibabaScbpTargetAdPlanListResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanListResponse struct {

    // 定向推广计划列表
    
    QuickCampaignList  struct {
        TopP4pBasicQuickCampaignView  []TopP4pBasicQuickCampaignView `json:"top_p4p_basic_quick_campaign_view,omitempty"`
    } `json:"quick_campaign_list,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanListResponse struct {

    // 定向推广计划列表
    QuickCampaignList   []TopP4pBasicQuickCampaignView `json:"quick_campaign_list,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

}
