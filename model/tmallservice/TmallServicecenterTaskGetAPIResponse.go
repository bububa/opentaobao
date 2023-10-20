package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskGetAPIResponse 服务工单拉取 API返回值
// tmall.servicecenter.task.get
//
// 接口供服务供应商通过交易主订单查询其未拉取的任务类工单
type TmallServicecenterTaskGetAPIResponse struct {
	model.CommonResponse
	TmallServicecenterTaskGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterTaskGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterTaskGetAPIResponseModel).Reset()
}

// TmallServicecenterTaskGetAPIResponseModel is 服务工单拉取 成功返回结果
type TmallServicecenterTaskGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_task_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServicePacket&lt;ServiceTaskDO&gt;
	ServiceTaskPacket *ServiceTaskPacket `json:"service_task_packet,omitempty" xml:"service_task_packet,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterTaskGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceTaskPacket = nil
}

var poolTmallServicecenterTaskGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterTaskGetAPIResponse)
	},
}

// GetTmallServicecenterTaskGetAPIResponse 从 sync.Pool 获取 TmallServicecenterTaskGetAPIResponse
func GetTmallServicecenterTaskGetAPIResponse() *TmallServicecenterTaskGetAPIResponse {
	return poolTmallServicecenterTaskGetAPIResponse.Get().(*TmallServicecenterTaskGetAPIResponse)
}

// ReleaseTmallServicecenterTaskGetAPIResponse 将 TmallServicecenterTaskGetAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterTaskGetAPIResponse(v *TmallServicecenterTaskGetAPIResponse) {
	v.Reset()
	poolTmallServicecenterTaskGetAPIResponse.Put(v)
}
