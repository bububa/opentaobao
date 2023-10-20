package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingVersionGenerateAPIResponse 生成发布使用的版本号 API返回值
// alibaba.wdk.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabaWdkMarketingVersionGenerateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingVersionGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingVersionGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingVersionGenerateAPIResponseModel).Reset()
}

// AlibabaWdkMarketingVersionGenerateAPIResponseModel is 生成发布使用的版本号 成功返回结果
type AlibabaWdkMarketingVersionGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_version_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingVersionGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingVersionGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingVersionGenerateAPIResponse)
	},
}

// GetAlibabaWdkMarketingVersionGenerateAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingVersionGenerateAPIResponse
func GetAlibabaWdkMarketingVersionGenerateAPIResponse() *AlibabaWdkMarketingVersionGenerateAPIResponse {
	return poolAlibabaWdkMarketingVersionGenerateAPIResponse.Get().(*AlibabaWdkMarketingVersionGenerateAPIResponse)
}

// ReleaseAlibabaWdkMarketingVersionGenerateAPIResponse 将 AlibabaWdkMarketingVersionGenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingVersionGenerateAPIResponse(v *AlibabaWdkMarketingVersionGenerateAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingVersionGenerateAPIResponse.Put(v)
}
