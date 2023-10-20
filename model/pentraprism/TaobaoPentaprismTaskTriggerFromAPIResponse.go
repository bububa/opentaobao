package pentraprism

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPentaprismTaskTriggerFromAPIResponse 任务进度推进（根据fromtoken） API返回值
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
type TaobaoPentaprismTaskTriggerFromAPIResponse struct {
	model.CommonResponse
	TaobaoPentaprismTaskTriggerFromAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPentaprismTaskTriggerFromAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPentaprismTaskTriggerFromAPIResponseModel).Reset()
}

// TaobaoPentaprismTaskTriggerFromAPIResponseModel is 任务进度推进（根据fromtoken） 成功返回结果
type TaobaoPentaprismTaskTriggerFromAPIResponseModel struct {
	XMLName xml.Name `xml:"pentaprism_task_trigger_from_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TOP接口标准出参
	Result *TaskResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPentaprismTaskTriggerFromAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPentaprismTaskTriggerFromAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPentaprismTaskTriggerFromAPIResponse)
	},
}

// GetTaobaoPentaprismTaskTriggerFromAPIResponse 从 sync.Pool 获取 TaobaoPentaprismTaskTriggerFromAPIResponse
func GetTaobaoPentaprismTaskTriggerFromAPIResponse() *TaobaoPentaprismTaskTriggerFromAPIResponse {
	return poolTaobaoPentaprismTaskTriggerFromAPIResponse.Get().(*TaobaoPentaprismTaskTriggerFromAPIResponse)
}

// ReleaseTaobaoPentaprismTaskTriggerFromAPIResponse 将 TaobaoPentaprismTaskTriggerFromAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPentaprismTaskTriggerFromAPIResponse(v *TaobaoPentaprismTaskTriggerFromAPIResponse) {
	v.Reset()
	poolTaobaoPentaprismTaskTriggerFromAPIResponse.Put(v)
}
