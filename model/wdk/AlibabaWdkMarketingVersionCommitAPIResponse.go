package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingVersionCommitAPIResponse 提交版本号 API返回值
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabaWdkMarketingVersionCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingVersionCommitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingVersionCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingVersionCommitAPIResponseModel).Reset()
}

// AlibabaWdkMarketingVersionCommitAPIResponseModel is 提交版本号 成功返回结果
type AlibabaWdkMarketingVersionCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_version_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingVersionCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingVersionCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingVersionCommitAPIResponse)
	},
}

// GetAlibabaWdkMarketingVersionCommitAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingVersionCommitAPIResponse
func GetAlibabaWdkMarketingVersionCommitAPIResponse() *AlibabaWdkMarketingVersionCommitAPIResponse {
	return poolAlibabaWdkMarketingVersionCommitAPIResponse.Get().(*AlibabaWdkMarketingVersionCommitAPIResponse)
}

// ReleaseAlibabaWdkMarketingVersionCommitAPIResponse 将 AlibabaWdkMarketingVersionCommitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingVersionCommitAPIResponse(v *AlibabaWdkMarketingVersionCommitAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingVersionCommitAPIResponse.Put(v)
}
