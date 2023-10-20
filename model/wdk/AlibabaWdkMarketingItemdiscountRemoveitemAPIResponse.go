package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse 移除报名的商品 API返回值
// alibaba.wdk.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
type AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItemdiscountRemoveitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItemdiscountRemoveitemAPIResponseModel is 移除报名的商品 成功返回结果
type AlibabaWdkMarketingItemdiscountRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse
func GetAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse() *AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse {
	return poolAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse.Get().(*AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse 将 AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse(v *AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountRemoveitemAPIResponse.Put(v)
}
