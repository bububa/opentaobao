package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeCreateactivityAPIResponse 创建全场活动 API返回值
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel is 创建全场活动 成功返回结果
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeCreateactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeCreateactivityAPIResponse
func GetAlibabaWdkMarketingFullrangeCreateactivityAPIResponse() *AlibabaWdkMarketingFullrangeCreateactivityAPIResponse {
	return poolAlibabaWdkMarketingFullrangeCreateactivityAPIResponse.Get().(*AlibabaWdkMarketingFullrangeCreateactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeCreateactivityAPIResponse 将 AlibabaWdkMarketingFullrangeCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeCreateactivityAPIResponse(v *AlibabaWdkMarketingFullrangeCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeCreateactivityAPIResponse.Put(v)
}
