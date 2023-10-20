package aetask

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AliexpressInteractiveTaskCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressInteractiveTaskCompleteAPIResponseModel).Reset()
}

// AliexpressInteractiveTaskCompleteAPIResponseModel is 任务完成接口 成功返回结果
type AliexpressInteractiveTaskCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_interactive_task_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressInteractiveTaskCompleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressInteractiveTaskCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressInteractiveTaskCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressInteractiveTaskCompleteAPIResponse)
	},
}

// GetAliexpressInteractiveTaskCompleteAPIResponse 从 sync.Pool 获取 AliexpressInteractiveTaskCompleteAPIResponse
func GetAliexpressInteractiveTaskCompleteAPIResponse() *AliexpressInteractiveTaskCompleteAPIResponse {
	return poolAliexpressInteractiveTaskCompleteAPIResponse.Get().(*AliexpressInteractiveTaskCompleteAPIResponse)
}

// ReleaseAliexpressInteractiveTaskCompleteAPIResponse 将 AliexpressInteractiveTaskCompleteAPIResponse 保存到 sync.Pool
func ReleaseAliexpressInteractiveTaskCompleteAPIResponse(v *AliexpressInteractiveTaskCompleteAPIResponse) {
	v.Reset()
	poolAliexpressInteractiveTaskCompleteAPIResponse.Put(v)
}
