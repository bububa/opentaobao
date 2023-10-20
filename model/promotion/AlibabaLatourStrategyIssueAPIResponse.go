package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLatourStrategyIssueAPIResponse 阿里巴巴权益发放接口 API返回值
// alibaba.latour.strategy.issue
//
// 阿里巴巴权益平台权益发放接口
type AlibabaLatourStrategyIssueAPIResponse struct {
	model.CommonResponse
	AlibabaLatourStrategyIssueAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLatourStrategyIssueAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLatourStrategyIssueAPIResponseModel).Reset()
}

// AlibabaLatourStrategyIssueAPIResponseModel is 阿里巴巴权益发放接口 成功返回结果
type AlibabaLatourStrategyIssueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_latour_strategy_issue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLatourStrategyIssueResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLatourStrategyIssueAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLatourStrategyIssueAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLatourStrategyIssueAPIResponse)
	},
}

// GetAlibabaLatourStrategyIssueAPIResponse 从 sync.Pool 获取 AlibabaLatourStrategyIssueAPIResponse
func GetAlibabaLatourStrategyIssueAPIResponse() *AlibabaLatourStrategyIssueAPIResponse {
	return poolAlibabaLatourStrategyIssueAPIResponse.Get().(*AlibabaLatourStrategyIssueAPIResponse)
}

// ReleaseAlibabaLatourStrategyIssueAPIResponse 将 AlibabaLatourStrategyIssueAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLatourStrategyIssueAPIResponse(v *AlibabaLatourStrategyIssueAPIResponse) {
	v.Reset()
	poolAlibabaLatourStrategyIssueAPIResponse.Put(v)
}
