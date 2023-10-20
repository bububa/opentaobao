package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse 删除买赠活动 API返回值
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponseModel is 删除买赠活动 成功返回结果
type AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse
func GetAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse() *AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse {
	return poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse.Get().(*AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse 将 AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse(v *AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse.Put(v)
}
