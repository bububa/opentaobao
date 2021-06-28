package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商分派工人 APIResponse
tmall.servicecenter.workcard.assignworker

服务商调用该接口分派工人给具体的工单
*/
type TmallServicecenterWorkcardAssignworkerAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardAssignworkerResponse `json:"tmall_servicecenter_workcard_assignworker_response,omitempty"` 
    TmallServicecenterWorkcardAssignworkerResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardAssignworkerResponse struct {

    // -
    
    Result  *struct {
        TmallServicecenterWorkcardAssignworkerResult  *TmallServicecenterWorkcardAssignworkerResult `json:"tmall_servicecenter_workcard_assignworker_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardAssignworkerResponse struct {

    // -
    Result   *TmallServicecenterWorkcardAssignworkerResult `json:"result,omitempty"`

}
