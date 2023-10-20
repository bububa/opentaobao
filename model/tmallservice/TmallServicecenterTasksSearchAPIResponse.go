package tmallservice

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallServicecenterTasksSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTasksSearchAPIResponseModel).Reset()
}

// TmallServicecenterTasksSearchAPIResponseModel is 查询任务类工单信息 成功返回结果
type TmallServicecenterTasksSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_tasks_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServicePacket&lt;ServiceTaskDO&gt;
	ServiceTaskPacket *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"service_task_packet,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTasksSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceTaskPacket = nil
}

var poolTmallServicecenterTasksSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTasksSearchAPIResponse)
	},
}

// GetTmallServicecenterTasksSearchAPIResponse 从 sync.Pool 获取 TmallServicecenterTasksSearchAPIResponse
func GetTmallServicecenterTasksSearchAPIResponse() *TmallServicecenterTasksSearchAPIResponse {
	return poolTmallServicecenterTasksSearchAPIResponse.Get().(*TmallServicecenterTasksSearchAPIResponse)
}

// ReleaseTmallServicecenterTasksSearchAPIResponse 将 TmallServicecenterTasksSearchAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTasksSearchAPIResponse(v *TmallServicecenterTasksSearchAPIResponse) {
	v.Reset()
	poolTmallServicecenterTasksSearchAPIResponse.Put(v)
}
