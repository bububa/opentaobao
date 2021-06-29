package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-获取计划的定向溢价数据 APIResponse
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据
*/
type AlibabaScbpTargetAdPlanTagGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanTagGetResponse
}

type AlibabaScbpTargetAdPlanTagGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_tag_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // TopP4pCampaignTargetingTagView
    
    Result   *TopP4pCampaignTargetingTagView `json:"result,omitempty" xml:"result,omitempty"`

    
}
