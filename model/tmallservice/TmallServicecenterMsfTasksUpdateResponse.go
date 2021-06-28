package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅工人任务批量完成接口 APIResponse
tmall.servicecenter.msf.tasks.update

喵师傅工人任务批量完成接口
*/
type TmallServicecenterMsfTasksUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterMsfTasksUpdateResponse `json:"tmall_servicecenter_msf_tasks_update_response,omitempty"` 
    TmallServicecenterMsfTasksUpdateResponse
}

/* model for simplify = false
type TmallServicecenterMsfTasksUpdateResponse struct {

    // result
    
    Result  *struct {
        ResultBase  *ResultBase `json:"result_base,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterMsfTasksUpdateResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
