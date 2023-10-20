package aetask

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressinteractivetaskcompleteAPIResponse 任务完成接口 API返回值
// aliexpress.interactive.task.complete
//
// 用户完成任务
type AliexpressinteractivetaskcompleteAPIResponse struct {
	model.CommonResponse
	AliexpressinteractivetaskcompleteAPIResponseModel
}

// AliexpressinteractivetaskcompleteAPIResponseModel is 任务完成接口 成功返回结果
type AliexpressinteractivetaskcompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_interactive_task_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressinteractivetaskcompleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
