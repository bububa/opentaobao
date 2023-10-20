package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeQueryitemAPIResponse 全场活动查询换购品 API返回值
// alibaba.wdk.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabaWdkMarketingFullrangeQueryitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeQueryitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel is 全场活动查询换购品 成功返回结果
type AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeQueryitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeQueryitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeQueryitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeQueryitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeQueryitemAPIResponse
func GetAlibabaWdkMarketingFullrangeQueryitemAPIResponse() *AlibabaWdkMarketingFullrangeQueryitemAPIResponse {
	return poolAlibabaWdkMarketingFullrangeQueryitemAPIResponse.Get().(*AlibabaWdkMarketingFullrangeQueryitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeQueryitemAPIResponse 将 AlibabaWdkMarketingFullrangeQueryitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeQueryitemAPIResponse(v *AlibabaWdkMarketingFullrangeQueryitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeQueryitemAPIResponse.Put(v)
}
