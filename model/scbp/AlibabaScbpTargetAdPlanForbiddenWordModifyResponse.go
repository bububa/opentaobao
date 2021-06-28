package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新增或删除屏蔽词 APIResponse
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词
*/
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanForbiddenWordModifyResponse `json:"alibaba_scbp_target_ad_plan_forbidden_word_modify_response,omitempty"` 
    AlibabaScbpTargetAdPlanForbiddenWordModifyResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanForbiddenWordModifyResponse struct {

    // true修改成功，fasle修改失败
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanForbiddenWordModifyResponse struct {

    // true修改成功，fasle修改失败
    Result   bool `json:"result,omitempty"`

}
