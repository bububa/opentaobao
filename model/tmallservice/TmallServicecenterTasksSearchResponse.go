package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单信息 APIResponse
tmall.servicecenter.tasks.search

查询任务类工单信息
*/
type TmallServicecenterTasksSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_tasks_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ServicePacket<ServiceTaskDO>
    
    ServiceTaskPacket   *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"