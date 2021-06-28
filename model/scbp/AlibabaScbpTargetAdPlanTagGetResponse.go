package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划的定向溢价数据 APIResponse
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据
*/
type AlibabaScbpTargetAdPlanTagGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanTagGetResponse `json:"alibaba_scbp_target_ad_plan_tag_get_response,omitempty"` 
    AlibabaScbpTargetAdPlanTagGetResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanTagGetResponse struct {

    // TopP4pCampaignTargetingTagView
    
    Result  *struct {
        TopP4pCampaignTargetingTagView  *TopP4pCampaignTargetingTagView `json:"top_p4p_campaign_targeting_tag_view,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanTagGetResponse struct {

    // TopP4pCampaignTargetingTagView
    Result   *TopP4pCampaignTargetingTagView `json:"result,omitempty"`

}
