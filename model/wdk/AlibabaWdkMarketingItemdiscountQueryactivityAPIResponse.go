package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse 查找特价活动 API返回值
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItemdiscountQueryactivityAPIResponseModel is 查找特价活动 成功返回结果
type AlibabaWdkMarketingItemdiscountQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询特价活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse
func GetAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse() *AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse {
	return poolAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse.Get().(*AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse 将 AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse(v *AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountQueryactivityAPIResponse.Put(v)
}
