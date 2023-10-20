package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterMsfTasksUpdateAPIResponse 喵师傅工人任务批量完成接口 API返回值
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
type TmallServicecenterMsfTasksUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterMsfTasksUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterMsfTasksUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterMsfTasksUpdateAPIResponseModel).Reset()
}

// TmallServicecenterMsfTasksUpdateAPIResponseModel is 喵师傅工人任务批量完成接口 成功返回结果
type TmallServicecenterMsfTasksUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_msf_tasks_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterMsfTasksUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterMsfTasksUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterMsfTasksUpdateAPIResponse)
	},
}

// GetTmallServicecenterMsfTasksUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterMsfTasksUpdateAPIResponse
func GetTmallServicecenterMsfTasksUpdateAPIResponse() *TmallServicecenterMsfTasksUpdateAPIResponse {
	return poolTmallServicecenterMsfTasksUpdateAPIResponse.Get().(*TmallServicecenterMsfTasksUpdateAPIResponse)
}

// ReleaseTmallServicecenterMsfTasksUpdateAPIResponse 将 TmallServicecenterMsfTasksUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterMsfTasksUpdateAPIResponse(v *TmallServicecenterMsfTasksUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterMsfTasksUpdateAPIResponse.Put(v)
}
