package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新增或删除屏蔽词 APIResponse
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词
*/
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_plan_forbidden_word_modify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true修改成功，fasle修改失败
    
    Result   bool `json:"result,omitempty" xml:"