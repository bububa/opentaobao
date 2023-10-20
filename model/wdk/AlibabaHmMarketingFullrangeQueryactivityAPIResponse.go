package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryactivityAPIResponse 全场活动查询活动 API返回值
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaHmMarketingFullrangeQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel is 全场活动查询活动 成功返回结果
type AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingFullrangeQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingFullrangeQueryactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingFullrangeQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingFullrangeQueryactivityAPIResponse
func GetAlibabaHmMarketingFullrangeQueryactivityAPIResponse() *AlibabaHmMarketingFullrangeQueryactivityAPIResponse {
	return poolAlibabaHmMarketingFullrangeQueryactivityAPIResponse.Get().(*AlibabaHmMarketingFullrangeQueryactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingFullrangeQueryactivityAPIResponse 将 AlibabaHmMarketingFullrangeQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingFullrangeQueryactivityAPIResponse(v *AlibabaHmMarketingFullrangeQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingFullrangeQueryactivityAPIResponse.Put(v)
}
