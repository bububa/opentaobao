package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse 查询特价商品 API返回值
// alibaba.wdk.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
type AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItemdiscountQueryitemsAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItemdiscountQueryitemsAPIResponseModel is 查询特价商品 成功返回结果
type AlibabaWdkMarketingItemdiscountQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse)
	},
}

// GetAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse
func GetAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse() *AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse {
	return poolAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse.Get().(*AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse)
}

// ReleaseAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse 将 AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse(v *AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountQueryitemsAPIResponse.Put(v)
}
