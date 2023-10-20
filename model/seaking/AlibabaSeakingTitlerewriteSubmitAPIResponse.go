package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTitlerewriteSubmitAPIResponse 提交标题改写任务 API返回值
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
type AlibabaSeakingTitlerewriteSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingTitlerewriteSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingTitlerewriteSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingTitlerewriteSubmitAPIResponseModel).Reset()
}

// AlibabaSeakingTitlerewriteSubmitAPIResponseModel is 提交标题改写任务 成功返回结果
type AlibabaSeakingTitlerewriteSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_titlerewrite_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingTitlerewriteSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaskId = 0
}

var poolAlibabaSeakingTitlerewriteSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingTitlerewriteSubmitAPIResponse)
	},
}

// GetAlibabaSeakingTitlerewriteSubmitAPIResponse 从 sync.Pool 获取 AlibabaSeakingTitlerewriteSubmitAPIResponse
func GetAlibabaSeakingTitlerewriteSubmitAPIResponse() *AlibabaSeakingTitlerewriteSubmitAPIResponse {
	return poolAlibabaSeakingTitlerewriteSubmitAPIResponse.Get().(*AlibabaSeakingTitlerewriteSubmitAPIResponse)
}

// ReleaseAlibabaSeakingTitlerewriteSubmitAPIResponse 将 AlibabaSeakingTitlerewriteSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingTitlerewriteSubmitAPIResponse(v *AlibabaSeakingTitlerewriteSubmitAPIResponse) {
	v.Reset()
	poolAlibabaSeakingTitlerewriteSubmitAPIResponse.Put(v)
}
