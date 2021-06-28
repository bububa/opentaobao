package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅工人任务批量完成接口 APIResponse
tmall.servicecenter.msf.tasks.update

喵师傅工人任务批量完成接口
*/
type TmallServicecenterMsfTasksUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_msf_tasks_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"