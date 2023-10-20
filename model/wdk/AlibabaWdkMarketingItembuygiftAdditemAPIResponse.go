package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftAdditemAPIResponse 增加买赠活动商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.wdk.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
type AlibabaWdkMarketingItembuygiftAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftAdditemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftAdditemAPIResponseModel is 增加买赠活动商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaWdkMarketingItembuygiftAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftAdditemAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftAdditemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftAdditemAPIResponse
func GetAlibabaWdkMarketingItembuygiftAdditemAPIResponse() *AlibabaWdkMarketingItembuygiftAdditemAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftAdditemAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftAdditemAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftAdditemAPIResponse 将 AlibabaWdkMarketingItembuygiftAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftAdditemAPIResponse(v *AlibabaWdkMarketingItembuygiftAdditemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftAdditemAPIResponse.Put(v)
}
