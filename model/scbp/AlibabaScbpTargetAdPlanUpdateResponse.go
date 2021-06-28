package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-更新推广计划的基础信息 APIResponse
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息
*/
type AlibabaScbpTargetAdPlanUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_plan_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true修改成功，false修改失败
    
    Result   bool `json:"result,omitempty" xml:"