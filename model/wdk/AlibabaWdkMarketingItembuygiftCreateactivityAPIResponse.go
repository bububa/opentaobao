package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse 创建买赠活动 API返回值
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftCreateactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftCreateactivityAPIResponseModel is 创建买赠活动 成功返回结果
type AlibabaWdkMarketingItembuygiftCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse
func GetAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse() *AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse 将 AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse(v *AlibabaWdkMarketingItembuygiftCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftCreateactivityAPIResponse.Put(v)
}
