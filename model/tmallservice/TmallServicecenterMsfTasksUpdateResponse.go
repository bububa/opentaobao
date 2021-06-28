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
    TmallServicecenterMsfTasksUpdateResponse
}

type TmallServicecenterMsfTasksUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_msf_tasks_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
