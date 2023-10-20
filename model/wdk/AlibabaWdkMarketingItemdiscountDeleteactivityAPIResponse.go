package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse 删除商品特价活动 API返回值
// alibaba.wdk.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponseModel is 删除商品特价活动 成功返回结果
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse
func GetAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse() *AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse {
	return poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse.Get().(*AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse 将 AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse(v *AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse.Put(v)
}
