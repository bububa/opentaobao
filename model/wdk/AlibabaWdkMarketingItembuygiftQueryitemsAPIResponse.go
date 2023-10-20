package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse 查询买赠活动下的商品 API返回值
// alibaba.wdk.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
type AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftQueryitemsAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftQueryitemsAPIResponseModel is 查询买赠活动下的商品 成功返回结果
type AlibabaWdkMarketingItembuygiftQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse
func GetAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse() *AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse 将 AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse(v *AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftQueryitemsAPIResponse.Put(v)
}
