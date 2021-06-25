package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单信息 APIResponse
tmall.servicecenter.tasks.search

查询任务类工单信息
*/
type TmallServicecenterTasksSearchAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterTasksSearchResponse `json:"tmall_servicecenter_tasks_search_response,omitempty"`
}

type TmallServicecenterTasksSearchResponse struct {

    // ServicePacket<ServiceTaskDO>
    ServiceTaskPacket   *ServiceTaskPacket `json:"service_task_packet,omitempty"`

}
