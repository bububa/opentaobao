package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskgetAPIResponse 服务工单拉取 API返回值
// tmall.servicecenter.task.get
//
// 接口供服务供应商通过交易主订单查询其未拉取的任务类工单
type TmallservicecentertaskgetAPIResponse struct {
	model.CommonResponse
	TmallservicecentertaskgetAPIResponseModel
}

// TmallservicecentertaskgetAPIResponseModel is 服务工单拉取 成功返回结果
type TmallservicecentertaskgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServicePacket&lt;ServiceTaskDO&gt;
	ServiceTaskPacket *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"service_task_packet,omitempty"`
}
