package aetask

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressInteractiveTaskCompleteAPIResponse 任务完成接口 API返回值
// aliexpress.interactive.task.complete
//
// 用户完成任务
type AliexpressInteractiveTaskCompleteAPIResponse struct {
	model.CommonResponse
	AliexpressInteractiveTaskCompleteAPIResponseModel
}

// AliexpressInteractiveTaskCompleteAPIResponseModel is 任务完成接口 成功返回结果
type AliexpressInteractiveTaskCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_interactive_task_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressInteractiveTaskCompleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
