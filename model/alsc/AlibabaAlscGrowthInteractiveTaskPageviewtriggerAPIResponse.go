package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse 浏览打点接口 API返回值
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
type AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel).Reset()
}

// AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel is 浏览打点接口 成功返回结果
type AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_pageviewtrigger_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打点返回
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse)
	},
}

// GetAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse
func GetAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse() *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse {
	return poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse.Get().(*AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse 将 AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse(v *AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponse.Put(v)
}
