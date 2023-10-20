package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountAdditemAPIResponse 报名特价商品 API返回值
// alibaba.wdk.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
type AlibabaWdkMarketingItemdiscountAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItemdiscountAdditemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItemdiscountAdditemAPIResponseModel is 报名特价商品 成功返回结果
type AlibabaWdkMarketingItemdiscountAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItemdiscountAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItemdiscountAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItemdiscountAdditemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItemdiscountAdditemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountAdditemAPIResponse
func GetAlibabaWdkMarketingItemdiscountAdditemAPIResponse() *AlibabaWdkMarketingItemdiscountAdditemAPIResponse {
	return poolAlibabaWdkMarketingItemdiscountAdditemAPIResponse.Get().(*AlibabaWdkMarketingItemdiscountAdditemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItemdiscountAdditemAPIResponse 将 AlibabaWdkMarketingItemdiscountAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountAdditemAPIResponse(v *AlibabaWdkMarketingItemdiscountAdditemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountAdditemAPIResponse.Put(v)
}
