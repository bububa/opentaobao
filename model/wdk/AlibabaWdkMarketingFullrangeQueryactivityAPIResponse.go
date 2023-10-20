package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeQueryactivityAPIResponse 全场活动查询活动 API返回值
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaWdkMarketingFullrangeQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeQueryactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeQueryactivityAPIResponseModel is 全场活动查询活动 成功返回结果
type AlibabaWdkMarketingFullrangeQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeQueryactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeQueryactivityAPIResponse
func GetAlibabaWdkMarketingFullrangeQueryactivityAPIResponse() *AlibabaWdkMarketingFullrangeQueryactivityAPIResponse {
	return poolAlibabaWdkMarketingFullrangeQueryactivityAPIResponse.Get().(*AlibabaWdkMarketingFullrangeQueryactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeQueryactivityAPIResponse 将 AlibabaWdkMarketingFullrangeQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeQueryactivityAPIResponse(v *AlibabaWdkMarketingFullrangeQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeQueryactivityAPIResponse.Put(v)
}
