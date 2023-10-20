package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse 任务领奖 API返回值
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
type AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponseModel).Reset()
}

// AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponseModel is 任务领奖 成功返回结果
type AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_receivetaskprize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发奖结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse)
	},
}

// GetAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse
func GetAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse() *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse {
	return poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse.Get().(*AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse)
}

// ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse 将 AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse(v *AlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveTaskReceivetaskprizeAPIResponse.Put(v)
}
