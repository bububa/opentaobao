package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeCreateactivityAPIResponse 创建全场活动 API返回值
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaHmMarketingFullrangeCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel is 创建全场活动 成功返回结果
type AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingFullrangeCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingFullrangeCreateactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingFullrangeCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingFullrangeCreateactivityAPIResponse
func GetAlibabaHmMarketingFullrangeCreateactivityAPIResponse() *AlibabaHmMarketingFullrangeCreateactivityAPIResponse {
	return poolAlibabaHmMarketingFullrangeCreateactivityAPIResponse.Get().(*AlibabaHmMarketingFullrangeCreateactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingFullrangeCreateactivityAPIResponse 将 AlibabaHmMarketingFullrangeCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeCreateactivityAPIResponse(v *AlibabaHmMarketingFullrangeCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeCreateactivityAPIResponse.Put(v)
}
