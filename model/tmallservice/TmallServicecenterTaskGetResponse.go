package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务工单拉取 APIResponse
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单
*/
type TmallServicecenterTaskGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterTaskGetResponse `json:"tmall_servicecenter_task_get_response,omitempty"` 
    TmallServicecenterTaskGetResponse
}

/* model for simplify = false
type TmallServicecenterTaskGetResponse struct {

    // ServicePacket<ServiceTaskDO>
    
    ServiceTaskPacket  *struct {
        ServiceTaskPacket  *ServiceTaskPacket `json:"service_task_packet,omitempty"`
    } `json:"service_task_packet,omitempty"`
    

}
*/

type TmallServicecenterTaskGetResponse struct {

    // ServicePacket<ServiceTaskDO>
    ServiceTaskPacket   *ServiceTaskPacket `json:"service_task_packet,omitempty"`

}
