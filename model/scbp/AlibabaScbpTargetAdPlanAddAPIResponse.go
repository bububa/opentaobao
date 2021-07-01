package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新建计划 API返回值 
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
type AlibabaScbpTargetAdPlanAddAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanAddAPIResponseModel
}

// 定向推广-新建计划 成功返回结果
type AlibabaScbpTargetAdPlanAddAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
