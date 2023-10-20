package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingVersionCommitAPIResponse 提交版本号 API返回值
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabaHmMarketingVersionCommitAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingVersionCommitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingVersionCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingVersionCommitAPIResponseModel).Reset()
}

// AlibabaHmMarketingVersionCommitAPIResponseModel is 提交版本号 成功返回结果
type AlibabaHmMarketingVersionCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_version_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingVersionCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingVersionCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingVersionCommitAPIResponse)
	},
}

// GetAlibabaHmMarketingVersionCommitAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingVersionCommitAPIResponse
func GetAlibabaHmMarketingVersionCommitAPIResponse() *AlibabaHmMarketingVersionCommitAPIResponse {
	return poolAlibabaHmMarketingVersionCommitAPIResponse.Get().(*AlibabaHmMarketingVersionCommitAPIResponse)
}

// ReleaseAlibabaHmMarketingVersionCommitAPIResponse 将 AlibabaHmMarketingVersionCommitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingVersionCommitAPIResponse(v *AlibabaHmMarketingVersionCommitAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingVersionCommitAPIResponse.Put(v)
}
