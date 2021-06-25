package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-更新推广计划的基础信息 APIResponse
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息
*/
type AlibabaScbpTargetAdPlanUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdPlanUpdateResponse `json:"alibaba_scbp_target_ad_plan_update_response,omitempty"`
}

type AlibabaScbpTargetAdPlanUpdateResponse struct {

    // true修改成功，false修改失败
    Result   bool `json:"result,omitempty"`

}
