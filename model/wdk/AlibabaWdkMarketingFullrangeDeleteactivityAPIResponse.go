package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse 全场活动删除活动接口 API返回值
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel is 全场活动删除活动接口 成功返回结果
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse
func GetAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse() *AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse {
	return poolAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse.Get().(*AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse 将 AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse(v *AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeDeleteactivityAPIResponse.Put(v)
}
