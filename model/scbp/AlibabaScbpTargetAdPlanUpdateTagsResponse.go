package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 APIResponse
alibaba.scbp.target.ad.plan.update.tags

定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
*/
type AlibabaScbpTargetAdPlanUpdateTagsAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanUpdateTagsResponse
}

type AlibabaScbpTargetAdPlanUpdateTagsResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_update_tags_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改记录数量
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
