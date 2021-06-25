package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新建计划 APIResponse
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
type AlibabaScbpTargetAdPlanAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdPlanAddResponse `json:"alibaba_scbp_target_ad_plan_add_response,omitempty"`
}

type AlibabaScbpTargetAdPlanAddResponse struct {

    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

}
