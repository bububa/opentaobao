package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTitlerewriteResultAPIResponse 获取标题改写任务结果 API返回值
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
type AlibabaSeakingTitlerewriteResultAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingTitlerewriteResultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingTitlerewriteResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingTitlerewriteResultAPIResponseModel).Reset()
}

// AlibabaSeakingTitlerewriteResultAPIResponseModel is 获取标题改写任务结果 成功返回结果
type AlibabaSeakingTitlerewriteResultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_titlerewrite_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingTitlerewriteResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSeakingTitlerewriteResultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingTitlerewriteResultAPIResponse)
	},
}

// GetAlibabaSeakingTitlerewriteResultAPIResponse 从 sync.Pool 获取 AlibabaSeakingTitlerewriteResultAPIResponse
func GetAlibabaSeakingTitlerewriteResultAPIResponse() *AlibabaSeakingTitlerewriteResultAPIResponse {
	return poolAlibabaSeakingTitlerewriteResultAPIResponse.Get().(*AlibabaSeakingTitlerewriteResultAPIResponse)
}

// ReleaseAlibabaSeakingTitlerewriteResultAPIResponse 将 AlibabaSeakingTitlerewriteResultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingTitlerewriteResultAPIResponse(v *AlibabaSeakingTitlerewriteResultAPIResponse) {
	v.Reset()
	poolAlibabaSeakingTitlerewriteResultAPIResponse.Put(v)
}
