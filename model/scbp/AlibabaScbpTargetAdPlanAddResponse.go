package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新建计划 APIResponse
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
type AlibabaScbpTargetAdPlanAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_plan_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 计划ID
    
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"