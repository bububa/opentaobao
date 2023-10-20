package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOnecodeIssueAPIResponse onecode代码通用鉴权 API返回值
// alibaba.interact.onecode.issue
//
// 手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
type AlibabaInteractOnecodeIssueAPIResponse struct {
	model.CommonResponse
	AlibabaInteractOnecodeIssueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractOnecodeIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractOnecodeIssueAPIResponseModel).Reset()
}

// AlibabaInteractOnecodeIssueAPIResponseModel is onecode代码通用鉴权 成功返回结果
type AlibabaInteractOnecodeIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_onecode_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractOnecodeIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractOnecodeIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractOnecodeIssueAPIResponse)
	},
}

// GetAlibabaInteractOnecodeIssueAPIResponse 从 sync.Pool 获取 AlibabaInteractOnecodeIssueAPIResponse
func GetAlibabaInteractOnecodeIssueAPIResponse() *AlibabaInteractOnecodeIssueAPIResponse {
	return poolAlibabaInteractOnecodeIssueAPIResponse.Get().(*AlibabaInteractOnecodeIssueAPIResponse)
}

// ReleaseAlibabaInteractOnecodeIssueAPIResponse 将 AlibabaInteractOnecodeIssueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractOnecodeIssueAPIResponse(v *AlibabaInteractOnecodeIssueAPIResponse) {
	v.Reset()
	poolAlibabaInteractOnecodeIssueAPIResponse.Put(v)
}
