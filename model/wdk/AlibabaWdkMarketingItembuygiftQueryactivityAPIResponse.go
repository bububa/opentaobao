package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse 查询买赠活动 API返回值
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel is 查询买赠活动 成功返回结果
type AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse
func GetAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse() *AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse 将 AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse(v *AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftQueryactivityAPIResponse.Put(v)
}
