package alsc

import (
	"encoding/xml"

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

// AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel is 浏览打点接口 成功返回结果
type AlibabaAlscGrowthInteractiveTaskPageviewtriggerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_pageviewtrigger_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打点返回
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
