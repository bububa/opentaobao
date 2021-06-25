package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单挂起 APIResponse
tmall.servicecenter.workcard.suspend

工单挂起
*/
type TmallServicecenterWorkcardSuspendAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardSuspendResponse `json:"tmall_servicecenter_workcard_suspend_response,omitempty"`
}

type TmallServicecenterWorkcardSuspendResponse struct {

    // 系统自动生成
    Result   *TmallServicecenterWorkcardSuspendResult `json:"result,omitempty"`

}
