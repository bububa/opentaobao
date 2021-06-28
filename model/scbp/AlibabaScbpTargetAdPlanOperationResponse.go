package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-计划开启/暂停/删除 APIResponse
alibaba.scbp.target.ad.plan.operation

定向推广-计划开启/暂停/删除
*/
type AlibabaScbpTargetAdPlanOperationAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpTargetAdPlanOperationResponse `json:"alibaba_scbp_target_ad_plan_operation_response,omitempty"` 
    AlibabaScbpTargetAdPlanOperationResponse
}

/* model for simplify = false
type AlibabaScbpTargetAdPlanOperationResponse struct {

    // 修改成功记录数
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpTargetAdPlanOperationResponse struct {

    // 修改成功记录数
    Result   int64 `json:"result,omitempty"`

}
