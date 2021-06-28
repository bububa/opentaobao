package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
预约失败 APIResponse
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
type TmallServicecenterWorkcardReservefailAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardReservefailResponse `json:"tmall_servicecenter_workcard_reservefail_response,omitempty"` 
    TmallServicecenterWorkcardReservefailResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardReservefailResponse struct {

    // -
    
    Result  *struct {
        TmallServicecenterWorkcardReservefailResult  *TmallServicecenterWorkcardReservefailResult `json:"tmall_servicecenter_workcard_reservefail_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardReservefailResponse struct {

    // -
    Result   *TmallServicecenterWorkcardReservefailResult `json:"result,omitempty"`

}
