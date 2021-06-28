package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单预约 APIResponse
tmall.servicecenter.workcard.reserve

服务工单更新通用接口
*/
type TmallServicecenterWorkcardReserveAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardReserveResponse `json:"tmall_servicecenter_workcard_reserve_response,omitempty"` 
    TmallServicecenterWorkcardReserveResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardReserveResponse struct {

    // 调用结果
    
    Result  *struct {
        TmallServicecenterWorkcardReserveResult  *TmallServicecenterWorkcardReserveResult `json:"tmall_servicecenter_workcard_reserve_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardReserveResponse struct {

    // 调用结果
    Result   *TmallServicecenterWorkcardReserveResult `json:"result,omitempty"`

}
