package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.wdk.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
type AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel is 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 移除商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse
func GetAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse() *AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse 将 AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse(v *AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftRemoveitemAPIResponse.Put(v)
}
