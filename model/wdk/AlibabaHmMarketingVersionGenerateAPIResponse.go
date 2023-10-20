package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingVersionGenerateAPIResponse 生成发布使用的版本号 API返回值
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabaHmMarketingVersionGenerateAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingVersionGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingVersionGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingVersionGenerateAPIResponseModel).Reset()
}

// AlibabaHmMarketingVersionGenerateAPIResponseModel is 生成发布使用的版本号 成功返回结果
type AlibabaHmMarketingVersionGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_version_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingVersionGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingVersionGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingVersionGenerateAPIResponse)
	},
}

// GetAlibabaHmMarketingVersionGenerateAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingVersionGenerateAPIResponse
func GetAlibabaHmMarketingVersionGenerateAPIResponse() *AlibabaHmMarketingVersionGenerateAPIResponse {
	return poolAlibabaHmMarketingVersionGenerateAPIResponse.Get().(*AlibabaHmMarketingVersionGenerateAPIResponse)
}

// ReleaseAlibabaHmMarketingVersionGenerateAPIResponse 将 AlibabaHmMarketingVersionGenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingVersionGenerateAPIResponse(v *AlibabaHmMarketingVersionGenerateAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingVersionGenerateAPIResponse.Put(v)
}
