package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务工单拉取 APIResponse
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单
*/
type TmallServicecenterTaskGetAPIResponse struct {
    model.CommonResponse
    TmallServicecenterTaskGetResponse
}

type TmallServicecenterTaskGetResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_task_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // ServicePacket<ServiceTaskDO>
    
    ServiceTaskPacket   *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"service_task_packet,omitempty"`

    
}
