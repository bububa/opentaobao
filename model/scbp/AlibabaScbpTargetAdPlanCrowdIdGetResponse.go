package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-人群标签ID获取(店铺老客、优选人群) APIResponse
alibaba.scbp.target.ad.plan.crowd.id.get

定向推广-人群标签ID获取(店铺老客、优选人群)
*/
type AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdPlanCrowdIdGetResponse `json:"alibaba_scbp_target_ad_plan_crowd_id_get_response,omitempty"`
}

type AlibabaScbpTargetAdPlanCrowdIdGetResponse struct {

    // 结果list
    ResultList   []CrowdView `json:"result_list,omitempty"`

}
