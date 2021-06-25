package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单改派门店 APIResponse
tmall.servicecenter.workcard.reassign

工单改派门店
*/
type TmallServicecenterWorkcardReassignAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardReassignResponse `json:"tmall_servicecenter_workcard_reassign_response,omitempty"`
}

type TmallServicecenterWorkcardReassignResponse struct {

    // 调用结果
    Result   *TmallServicecenterWorkcardReassignResult `json:"result,omitempty"`

}
