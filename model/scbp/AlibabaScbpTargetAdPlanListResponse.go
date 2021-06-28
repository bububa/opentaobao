package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-查询定向推广计划列表并返回计划基础信息 APIResponse
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息
*/
type AlibabaScbpTargetAdPlanListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_plan_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 定向推广计划列表
    
    QuickCampaignList   []TopP4pBasicQuickCampaignView `json:"quick_campaign_list,omitempty" xml:"