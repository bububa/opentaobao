package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse 领取任务 API返回值
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
type AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel).Reset()
}

// AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel is 领取任务 成功返回结果
type AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_receivetask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse)
	},
}

// GetAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse
func GetAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse() *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse {
	return poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse.Get().(*AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse 将 AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse(v *AlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskReceivetaskAPIResponse.Put(v)
}
