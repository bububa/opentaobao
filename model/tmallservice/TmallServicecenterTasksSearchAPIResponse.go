package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTasksSearchAPIResponse 查询任务类工单信息 API返回值
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
type TmallServicecenterTasksSearchAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTasksSearchAPIResponseModel
}

// TmallServicecenterTasksSearchAPIResponseModel is 查询任务类工单信息 成功返回结果
type TmallServicecenterTasksSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tasks_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServicePacket<ServiceTaskDO>
	ServiceTaskPacket *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"service_task_packet,omitempty"`
}
