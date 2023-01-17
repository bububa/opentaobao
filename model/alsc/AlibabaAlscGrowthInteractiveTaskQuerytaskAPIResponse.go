package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponse 本地生活用户增长互动任务平台查询任务 API返回值
// alibaba.alsc.growth.interactive.task.querytask
//
// 本地生活用户增长互动任务平台查询任务
type AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponseModel
}

// AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponseModel is 本地生活用户增长互动任务平台查询任务 成功返回结果
type AlibabaAlscGrowthInteractiveTaskQuerytaskAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_task_querytask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
